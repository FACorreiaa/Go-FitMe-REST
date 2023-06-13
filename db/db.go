package db

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresInstance struct {
	db *sqlx.DB
}

var DB *sqlx.DB

type QueryExecMode uint

const (
	CacheStatement = iota
)

func (m QueryExecMode) value() string {
	switch m {
	case CacheStatement:
		return "cache_statement"
	default:
		return ""
	}
}

type Config struct {
	host     string
	port     string
	username string
	password string
	dbName   string
	sslMode  string
	//maxConnWaitingTime   time.Duration
	//defaultQueryExecMode QueryExecMode
}

func SetupDatabase() (*sqlx.DB, error) {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGREST_PORT")
	sslMode := os.Getenv("POSTGRES_SSLMode")
	db, err := sqlx.Open(
		"pgx",
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v?sslmode=%v",
			username,
			password,
			dbHost,
			dbPort,
			dbName,
			sslMode,
		),
	)

	if err != nil {
		return db, nil
	}

	fmt.Println("Connection Opened to Database")
	DB = db
	//pingCtx, cancel := context.WithTimeout(context.Background(), configuration.maxConnWaitingTime)
	//defer cancel()
	//if err := db.Ping(pingCtx); err != nil {
	//	return nil, err
	//}
	return db, nil
}

func (p *PostgresInstance) GetDB() *sqlx.DB {
	return p.db
}
