package internals

import (
	_ "github.com/jackc/pgx/v5/stdlib"
)

//type Postgres struct {
//	db *sqlx.DB
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
//type Config struct {
//	host                 string
//	port                 string
//	username             string
//	password             string
//	dbName               string
//	sslMode              string
//	maxConnWaitingTime   time.Duration
//	defaultQueryExecMode QueryExecMode
//}
//
//func NewConfig(
//	host string,
//	port string,
//	username string,
//	password string,
//	dbName string,
//	sslMode string,
//	maxConnWaitingTime time.Duration,
//	defaultQueryExecMode QueryExecMode,
//) (Config, error) {
//	return Config{
//		host:                 host,
//		port:                 port,
//		username:             username,
//		password:             password,
//		dbName:               dbName,
//		sslMode:              sslMode,
//		maxConnWaitingTime:   maxConnWaitingTime,
//		defaultQueryExecMode: defaultQueryExecMode,
//	}, nil
//}
//
//func newDB(configuration Config) (*sqlx.DB, error) {
//	db, err := sqlx.Open(
//		"pgx",
//		fmt.Sprintf(
//			"postgres://%v:%v@%v:%v/%v?sslmode=%v&default_query_exec_mode=%v",
//			configuration.username,
//			configuration.password,
//			configuration.host,
//			configuration.port,
//			configuration.dbName,
//			configuration.sslMode,
//			configuration.defaultQueryExecMode.value(),
//		),
//	)
//
//	if err != nil {
//		return nil, err
//	}
//	_, cancel := context.WithTimeout(context.Background(), configuration.maxConnWaitingTime)
//	defer cancel()
//	if err := db.Ping(); err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}
//
//func NewPostgres(configuration Config) *Postgres {
//	db, err := newDB(configuration)
//	if err != nil {
//		logs.DefaultLogger.WithError(err).Fatal("Error on postgres init")
//		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
//	}
//	return &Postgres{db: db}
//}
//
//func (p *Postgres) GetDB() *sqlx.DB {
//	return p.db
//}
