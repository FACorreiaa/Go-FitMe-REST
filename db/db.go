package db

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
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
	host                 string
	port                 string
	username             string
	password             string
	dbName               string
	sslMode              string
	maxConnWaitingTime   time.Duration
	defaultQueryExecMode QueryExecMode
}

func SetupDatabase() error {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGREST_PORT")
	sslMode := os.Getenv("POSTGRES_SSLMode")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		username,
		password,
		dbName,
		dbHost,
		dbPort,
		sslMode)
	db, err := sqlx.Open("pgx", dsn)
	println("db")

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB = db
	//pingCtx, cancel := context.WithTimeout(context.Background(), config.maxConnWaitingTime)
	//defer cancel()
	//if err := db.Ping(pingCtx); err != nil {
	//	return nil, err
	//}
	return nil
}

func (p *PostgresInstance) GetDB() *sqlx.DB {
	return p.db
}
