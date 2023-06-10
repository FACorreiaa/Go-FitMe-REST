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
