package user

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

//func (u *userRepository) Migrate() error {
//	return u.DB.AutoMigrate(&domain.User{}).Error
//}

func (s *Repository) CreateNewUser(data NewUser) (int, error) {
	var id int
	query := `INSERT INTO users (username, email, password)
				VALUES (:username, :email, :password)
				RETURNING id`
	namedStmt, err := s.db.PrepareNamed(query)
	if err != nil {
		return 0, fmt.Errorf("error preparing named statement: %w", err)
	}

	err = namedStmt.QueryRow(data).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting values: %w", err)
	}

	return id, nil
}
