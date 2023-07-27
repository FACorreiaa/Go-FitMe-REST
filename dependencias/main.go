package dependencias

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/user"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthDependencies interface {
	GetDB() *sqlx.DB
	GetRedisClient() *redis.Client
}
type Dependencies interface {
	GetActivityService() *activity.ServiceActivity
	GetUserService() *user.ServiceUser
}