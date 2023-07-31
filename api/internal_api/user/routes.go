package user

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
)

func RoutesUser(deps DependenciesUser, sessionManager *auth.SessionManager) *chi.Mux {
	u := NewUserHandler(deps, sessionManager)

	router := chi.NewRouter()

	router.Post("/sign-up", u.SignUpUser)
	router.Post("/sign-in", u.SignInUser)
	router.Post("/sign-out", u.SignOutUser)
	router.Get("/user/info", u.GetUserInfo)

	return router
}
