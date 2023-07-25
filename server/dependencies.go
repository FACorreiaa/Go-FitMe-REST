package server

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AppDependencies struct {
	ActivityService *activity.ServiceActivity
	UserService     *user.ServiceUser
	// You may add other dependencies here if needed
}

func (ad *AppDependencies) GetActivityService() *activity.ServiceActivity {
	return ad.ActivityService
}

func (ad *AppDependencies) GetUserService() *user.ServiceUser {
	return ad.UserService
}

type SessionDependencies struct {
	DB  *sqlx.DB
	RDB *redis.Client
	// You may add other dependencies here if needed
}

func (ad *SessionDependencies) GetDB() *sqlx.DB {
	return ad.DB
}

func (ad *SessionDependencies) GetRedisClient() *redis.Client {
	return ad.RDB
}
