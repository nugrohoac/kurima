package kurima

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// BaseStruct .
type BaseStruct struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	CreatedBy string    `json:"-"`
	UpdatedBy string    `json:"-"`
}

// User .
type User struct {
	ID       string   `json:"id"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required"`
	Role     []string `json:"role"`
	Status   string   `json:"status"`
	BaseStruct
}

// Claim .
type Claim struct {
	User
	jwt.StandardClaims
}
