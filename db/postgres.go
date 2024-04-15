package db

import (
	"context"
	"crypto/md5"
	"embed"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

type DatabaseConfig struct {
	ConnectionURL string
}

//go:embed migrations/*.sql
var migrationFS embed.FS

const retries = 25

func NewDatabaseConfig() (*DatabaseConfig, error) {
	// TODO CONFIG VIPER HERE
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "dev" {
		if err != nil {
			log.Println(err)
			log.Fatal("Error loading .env file")
		}
	}

	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, errors.New("invalid DB_PORT")
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	schema := os.Getenv("")

	query := url.Values{
		"sslmode":  []string{"disable"},
		"timezone": []string{"utc"},
	}
	if schema != "" {
		query.Add("search_path", schema)
	}
	connURL := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, pass),
		Host:     host + ":" + strconv.Itoa(port),
		Path:     dbname,
		RawQuery: query.Encode(),
	}
	return &DatabaseConfig{
		ConnectionURL: connURL.String(),
	}, nil
}

// Init Init.
func Init(connectionURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		return nil, err
	}
	cfg.AfterConnect = func(_ context.Context, conn *pgx.Conn) error {
		uuid.Register(conn.TypeMap())
		return nil
	}

	return pgxpool.NewWithConfig(context.Background(), cfg)
}

func InitRedis(host, password string, db int) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	}), nil
}

func Migrate(conn *pgxpool.Pool) error {
	// migrate db
	slog.Info("Running migrations")
	ctx := context.Background()
	files, err := migrationFS.ReadDir("migrations")
	if err != nil {
		return err
	}

	slog.Info("Creating migrations table")
	_, err = conn.Exec(ctx, `
		create table if not exists _migrations (
			name text primary key,
			hash text not null,
			created_at timestamp default now()
		);
	`)
	if err != nil {
		return err
	}

	slog.Info("Checking applied migrations")
	rows, _ := conn.Query(ctx, `select name, hash from _migrations order by created_at desc`)
	var name, hash string
	appliedMigrations := make(map[string]string)
	_, err = pgx.ForEachRow(rows, []any{&name, &hash}, func() error {
		appliedMigrations[name] = hash
		return nil
	})

	if err != nil {
		return err
	}

	for _, file := range files {
		contents, err := migrationFS.ReadFile("migrations/" + file.Name())
		if err != nil {
			return err
		}

		contentHash := fmt.Sprintf("%x", md5.Sum(contents))

		if prevHash, ok := appliedMigrations[file.Name()]; ok {
			if prevHash != contentHash {
				return errors.New("hash mismatch for")
			}

			slog.Info(file.Name() + " already applied")
			continue
		}

		err = pgx.BeginFunc(ctx, conn, func(tx pgx.Tx) error {
			if _, err = tx.Exec(ctx, string(contents)); err != nil {
				return err
			}

			if _, err := tx.Exec(ctx, `insert into _migrations (name, hash) values ($1, $2)`,
				file.Name(), contentHash); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}
		slog.Info(file.Name() + " applied")
	}

	slog.Info("Migrations finished")
	return nil
}

// WaitForDB Small hack to wait for database to start inside docker.
func WaitForDB(pgpool *pgxpool.Pool) {
	ctx := context.Background()

	for attempts := 1; ; attempts++ {
		if attempts > retries {
			break
		}

		if err := pgpool.Ping(ctx); err == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
	}
}
