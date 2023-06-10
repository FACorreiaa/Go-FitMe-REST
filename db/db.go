//package db
//
//import (
//	"context"
//	"fmt"
//	"github.com/jackc/pgx/v5"
//	"github.com/jackc/pgx/v5/pgxpool"
//	"log"
//	"os"
//	"strconv"
//	"time"
//)
//
//type Postgres struct {
//	db *pgxpool.Pool
//}
//
//type QueryExecMode uint
//
//const (
//	CacheStatement = iota
//)
//
//func (m QueryExecMode) value() string {
//	switch m {
//	case CacheStatement:
//		return "cache_statement"
//	default:
//		return ""
//	}
//}
//
//type DefaultModel struct {
//	ID        uint       `gorm:"primary_key" json:"id"`
//	CreatedAt time.Time  `json:"createdAt"`
//	UpdatedAt time.Time  `json:"updatedAt"`
//	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
//}
//
//func SetupDatabase() {
//	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
//	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
//	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"

//	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%v", username, password, dbHost, dbName, sslMode)
//
//	conn, err := pgx.Connect(context.Background(), dsn)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//
//	defer conn.Close(context.Background())
//
//	fmt.Println("Connection Opened to Database")
//}

package db

import (
	"context"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/handlers/logs"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"syscall"
	"time"
)

type Postgres struct {
	db *pgxpool.Pool
}

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

//func NewConfig(
//	host string,
//	port string,
//	username string,
//	password string,
//	dbName string,
//	sslMode string,
//	maxConnWaitingTime time.Duration,
//	defaultQueryExecMode QueryExecMode,
//) Config {
//	return Config{
//		host:                 host,
//		port:                 port,
//		username:             username,
//		password:             password,
//		dbName:               dbName,
//		sslMode:              sslMode,
//		maxConnWaitingTime:   maxConnWaitingTime,
//		defaultQueryExecMode: defaultQueryExecMode,
//	}
//}

func newDB(config Config) (*pgxpool.Pool, error) {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGREST_PORT")
	sslMode := os.Getenv("POSTGRES_SSLMode")

	db, err := pgxpool.New(
		context.TODO(),
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v?sslmode=%v&default_query_exec_mode=%v",
			username,
			password,
			dbHost,
			dbPort,
			dbName,
			sslMode,
			config.defaultQueryExecMode.value(),
		),
	)

	if err != nil {
		return nil, err
	}
	//pingCtx, cancel := context.WithTimeout(context.Background(), config.maxConnWaitingTime)
	//defer cancel()
	//if err := db.Ping(pingCtx); err != nil {
	//	return nil, err
	//}

	return db, nil
}

func NewPostgres(config Config) *Postgres {
	db, err := newDB(config)
	if err != nil {
		logs.DefaultLogger.WithError(err).Fatal("Error on postgres init")
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}
	return &Postgres{db: db}
}

func (p *Postgres) GetDB() *pgxpool.Pool {
	return p.db
}
