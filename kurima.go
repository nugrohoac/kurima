package kurima

import "context"

// UserService .
type UserService interface {
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, usr User) (User, error)
	UpdatePassword(ctx context.Context, ID string, user User) (User, error)
}

// UserRepository .
type UserRepository interface {
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, user User) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByID(ctx context.Context, ID string) (User, error)
	UpdatePassword(ctx context.Context, ID string, user User) (User, error)
}

// JWTHash .
type JWTHash interface {
	Encode(user User) (string, error)
	Decode(tokenString string, claim *Claim) (bool, error)
}

// BcryptHash is interface to handle salt using bcrypt hash
type BcryptHash interface {
	Generate(password string) (string, error)
	Compare(hashedPassword, password string) error
}
