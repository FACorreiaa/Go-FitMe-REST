package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func RoutesUser(lg *logrus.Logger, db *sqlx.DB) *chi.Mux {
	u := NewUserHandler(lg, db)

	router := chi.NewRouter()

	router.Post("/sign-up", u.SignUpUser)
	router.Post("/sign-in", u.SignInUser)
	router.Get("/me", u.GetUserInfo)
	router.Post("/sign-out", u.SignOutUser)

	return router
}
