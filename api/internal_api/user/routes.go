package user

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/dependencies"
	"github.com/go-chi/chi/v5"
)

func RoutesUser(deps dependencies.Dependencies, sessionManager *auth.SessionManager) *chi.Mux {
	u := NewUserHandler(deps, sessionManager)

	router := chi.NewRouter()

	router.Post("/sign-up", u.SignUpUser)
	router.Post("/sign-in", u.SignInUser)
	router.Get("/me", u.GetUserInfo)
	router.Post("/sign-out", u.SignOutUser)

	return router
}
