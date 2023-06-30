package user

import "time"

type RoleType string

const (
	User  RoleType = "USER"
	Admin RoleType = "ADMIN"
)

type NewUser struct {
	ID        int        `json:"id,string" pg:"default:gen_random_uuid()"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      RoleType   `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
