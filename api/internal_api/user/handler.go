package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

type DependenciesUser interface {
	GetUserService() *ServiceUser
}

type Handler struct {
	router         *chi.Router
	dependencies   DependenciesUser
	ctx            context.Context
	sessionManager *auth.SessionManager
}

func NewUserHandler(deps DependenciesUser, sessionManager *auth.SessionManager) Handler {
	return Handler{
		dependencies:   deps,
		sessionManager: sessionManager,
	}
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

type userRequestBody struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required,min=6,max=48"`
}

type signUpSuccessResponse struct {
	ID int `json:"id"`
}

type signInRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// SignOutUser Sign out current user
//
// @Summary	Sign out a user
// @Tags		users
// @Accept		json
// @Produce	json
// @Success	200	{object}	SuccessResponse
// @Router		/users/sign-out [post]
func (u Handler) SignOutUser(w http.ResponseWriter, r *http.Request) {
	sessionHeader := r.Header.Get("Authorization")

	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		log.Printf("invalid session header")
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	sessionId := sessionHeader[7:]

	err := u.sessionManager.SignOut(sessionId)
	if err != nil {
		log.Printf("Error signing out: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetUserInfo godoc
// @Summary Get the user's info
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} auth.UserSession
// @Router /users/me [get]
func (u Handler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	sessionHeader := r.Header.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		log.Printf("invalid session header")
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	sessionId := sessionHeader[7:]

	user, err := u.sessionManager.GetSession(sessionId)
	if err != nil {
		log.Printf("error getting session")
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

// SignInUser godoc
// @Summary Sign in a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body signInRequestBody true "The user's email and password"
// @Success 200 {object} SuccessResponse
// @Header 200 {string} Authorization "contains the session id in bearer format"
// @Router /users/sign-in [post]
func (u Handler) SignInUser(w http.ResponseWriter, r *http.Request) {
	var user signInRequestBody

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse the request body
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Failed to parse request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validate the user struct
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		log.Println("Failed to validate user fields", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	sessionId, err := u.sessionManager.SignIn(user.Email, user.Password)
	if err != nil {
		log.Println("Failed to sign in user", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))

	// Send a response
	//r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User logged successfully"))

}

// SignUpUser godoc
// @Summary Sign up a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body userRequestBody true "The user's first name, last name, email, and password"
// @Success 200 {object} signUpSuccessResponse
// @Header 200 {string} Authorization "contains the session id in bearer format"
// @Router /users/sign-up [post]
func (u Handler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user userRequestBody
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse the request body
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Failed to parse request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash the password", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// create the user
	userID, err := u.dependencies.GetUserService().Create(NewUser{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		log.Println("Failed to create new user:", err)
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		return
	}
	sessionId, err := u.sessionManager.GenerateSession(auth.UserSession{
		Id:       userID,
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		log.Println("Failed to get session:", err)
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))
	// Create the response payload
	res := signUpSuccessResponse{
		ID: userID,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	// Send the response
	//w.Header().Set("Bearer", sessionId)
}
