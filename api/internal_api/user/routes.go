package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RoutesUser(db *sqlx.DB) *chi.Mux {
	u := NewUserHandler(db)

	router := chi.NewRouter()

	router.Post("/sign-up", u.SignUpUser)
	router.Post("/sign-in", u.SignInUser)
	router.Get("/me", u.GetUserInfo)
	router.Post("/sign-out", u.SignOutUser)

	return router
}
