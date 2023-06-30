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
	SslMode  string
}

func Connect(cnf ConfigDB) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cnf.Host,
		cnf.Port,
		cnf.User,
		cnf.Password,
		cnf.Name,
		cnf.SslMode,
	)
	db, err := sqlx.Connect("pgx", dsn)
	return db, err
}
