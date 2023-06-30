package user

//import (
//	"github.com/jackc/pgx/v5/pgxpool"
//)
//
//type userRepository struct {
//	db *pgxpool.Pool
//}
//
//// NewUserRepository : get injected database
//func NewUserRepository(db *pgxpool.Pool) RepositoryUser {
//	return &userRepository{
//		db: db,
//	}
//}
//func (u *userRepository) Save(user *User) (*User, error) {
//	return user, u.DB.Create(user).Error
//}
//func (u *userRepository) FindAll() ([]domain.User, error) {
//	var users []User
//	err := u.DB.Find(&users).Error
//	return users, err
//}
//func (u *userRepository) Delete(user *User) error {
//	return u.DB.Delete(&user).Error
//}
//func (u *userRepository) Migrate() error {
//	return u.DB.AutoMigrate(&domain.User{}).Error
//}

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) Repository {
	return Repository{db: db}
}

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
