package db

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type ConfigDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMODE  string
}

func Connect(cnf ConfigDB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s SSLMODE=%s",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
		cnf.SSLMODE,
	)
	db, err := sqlx.Connect("pgx", dsn)
	return db, err
}
