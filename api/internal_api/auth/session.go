package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type SessionManager struct {
	Rdb *redis.Client
	DB  *sqlx.DB
}

func NewSessionManager(rdb *redis.Client, db *sqlx.DB) *SessionManager {
	return &SessionManager{Rdb: rdb, DB: db}
}

type UserSession struct {
	Id       int
	Username string
	Email    string
}

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func validateSessionHeader(sessionHeader string) (string, error) {
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		return "", fmt.Errorf("invalid session header")
	}

	return sessionHeader[7:], nil
}

func GetSessionID(w http.ResponseWriter, r *http.Request) (string, error) {
	sessionHeader := r.Header.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	sessionId, err := validateSessionHeader(sessionHeader)

	if err != nil {
		log.Printf("invalid session header")
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return "", err
	}

	sessionId = sessionHeader[7:]

	return sessionId, nil
}

func (s *SessionManager) GenerateSession(data UserSession) (string, error) {
	sessionId := uuid.NewString()
	jsonData, _ := json.Marshal(data)
	err := s.Rdb.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	return sessionId, nil
}

func (s *SessionManager) SignIn(email, password string) (string, error) {
	// check if the user exists
	var user User
	err := s.DB.QueryRow("select id, username, email, password from users where email = $1", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return "", err
	}

	// check if the password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// create the session
	sessionId := uuid.NewString()
	jsonData, _ := json.Marshal(UserSession{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	})
	err = s.Rdb.Set(context.Background(), sessionId, string(jsonData), 24*time.Hour).Err()
	if err != nil {
		return "", err
	}
	println(sessionId)
	return sessionId, nil
}

func (s *SessionManager) SignOut(sessionId string) error {
	return s.Rdb.Del(context.Background(), sessionId).Err()
}

func (s *SessionManager) GetSession(session string) (*UserSession, error) {
	data, err := s.Rdb.Get(context.Background(), session).Result()
	if err != nil {
		return nil, err
	}

	// unmarshal the data
	var userSession UserSession
	err = json.Unmarshal([]byte(data), &userSession)
	if err != nil {
		return nil, err
	}

	return &userSession, nil
}

func SessionMiddleware(sessionManager *SessionManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Initialize the session manager and add it to the request context
			println(sessionManager)
			ctx := context.WithValue(r.Context(), "sessionManager", sessionManager)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
