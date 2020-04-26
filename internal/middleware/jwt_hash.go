package middleware

import (
	"github.com/nac-project/kurima"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/pkg/errors"
)

type jwtToken struct {
	secretKey []byte
	duration  time.Duration
}

// Encode .
func (j jwtToken) Encode(user kurima.User) (string, error) {
	expireTime := time.Now().Add(j.duration)

	claim := kurima.Claim{
		User: kurima.User{
			Email:  user.Email,
			Role:   user.Role,
			Status: user.Status,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", errors.Wrap(err, "error create token jwt")
	}

	return tokenString, nil
}

// Decode .
func (j jwtToken) Decode(tokenString string, claim *kurima.Claim) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return false, errors.Wrap(err, "error parse jwt")
	}

	if !token.Valid {
		return false, errors.Wrap(errors.New("error validation token"), "error validation token")
	}

	return true, nil
}

// NewJWTHash .
func NewJWTHash(secretKey []byte, duration time.Duration) kurima.JWTHash {
	return jwtToken{
		secretKey: secretKey,
		duration:  duration,
	}
}
