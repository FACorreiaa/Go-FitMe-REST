package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

type UserHandler struct {
	logger         *logrus.Logger
	router         *chi.Router
	userService    Service
	ctx            context.Context
	sessionManager *auth.SessionManager
}

func NewUserHandler(lg *logrus.Logger, db *sqlx.DB, sessionManager *auth.SessionManager) UserHandler {
	return UserHandler{
		logger:         lg,
		userService:    NewService(NewUserRepository(db)),
		sessionManager: sessionManager,
	}
}

//import (
//	"net/http"
//)
//
//type userController struct {
//	userService domain.UserService
//	fbService   service.FirebaseService
//}
//type UserController interface {
//	GetUsers(c *fiber.Context)
//	AddUser(c *gin.Context)
//}
//
//// NewUserController: constructor, dependency injection from user service and firebase service
//func NewUserController(s domain.UserService, f service.FirebaseService) UserController {
//	return &userController{
//		userService: s,
//		fbService:   f,
//	}
//}
//func (u *userController) GetUsers(c *gin.Context) {
//	users, err := u.userService.FindAll()
//	if err != nil {
//		sentry.CaptureException(err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
//		return
//	}
//	c.JSON(http.StatusOK, users)
//}
//func (u *userController) AddUser(c *gin.Context) {
//	var user domain.User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		sentry.CaptureException(err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	if err1 := (u.userService.Validate(&user)); err1 != nil {
//		sentry.CaptureException(err1)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
//		return
//	}
//	if ageValidation := (u.userService.ValidateAge(&user)); ageValidation != true {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB"})
//		return
//	}
//	uid, err := u.fbService.CreateUser(user.Email, user.Password)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldnot create user in firebase"})
//		return
//	}
//	user.ID = uid
//	u.userService.Create(&user)
//	c.JSON(http.StatusOK, user)
//}

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
func (u UserHandler) SignOutUser(w http.ResponseWriter, r *http.Request) {
	// get the session from the authorization header
	sessionHeader := r.Header.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		log.Printf("invalid session header")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// get the session id
	sessionId := sessionHeader[7:]

	// delete the session
	err := u.sessionManager.SignOut(sessionId)
	if err != nil {
		log.Printf("Error signing out: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
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
func (u UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	sessionHeader := r.Header.Get("Authorization")

	// ensure the session header is not empty and in the correct format
	if sessionHeader == "" || len(sessionHeader) < 8 || sessionHeader[:7] != "Bearer " {
		log.Printf("invalid session header")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	sessionId := sessionHeader[7:]

	user, err := u.sessionManager.GetSession(sessionId)
	if err != nil {
		log.Printf("error getting session")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
func (u UserHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
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

	// Send a response
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))
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
func (u UserHandler) SignUpUser(w http.ResponseWriter, r *http.Request) {
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
	userID, err := u.userService.Create(NewUser{
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

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", sessionId))

	// Create the response payload
	res := signUpSuccessResponse{
		ID: userID,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Println("Failed to encode response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
