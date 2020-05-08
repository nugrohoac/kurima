package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	_errors "github.com/pkg/errors"

	"github.com/nac-project/kurima"
)

// Middleware .
type Middleware struct {
	jwtHash  kurima.JWTHash
	userRepo kurima.UserRepository
}

// Auth .
func (m Middleware) Auth(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var claim kurima.Claim
		ctx := c.Request().Context()
		token := c.Request().Header.Get(echo.HeaderAuthorization)
		tokens := strings.Split(token, " ")
		if tokens[0] != "Bearer" {
			return errors.New("invalid token")
		}

		valid, err := m.jwtHash.Decode(tokens[1], &claim)
		if err != nil {
			return err
		}

		if !valid {
			return kurima.ErrInValidUser
		}

		if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) < 30*time.Second {
			user, err := m.userRepo.GetByEmail(ctx, claim.Email)
			if err != nil {
				return err
			}

			token, err = m.jwtHash.Encode(user)
			if err != nil {
				return err
			}

			token += "Bearer "
		}

		ctx = kurima.SetClaimOnContext(ctx, claim)
		ctx = kurima.SetTokenOnContext(ctx, token)
		c.SetRequest(c.Request().WithContext(ctx))

		return handler(c)
	}
}

// ErrHandler .
func (m Middleware) ErrHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := handler(c)
		if err != nil {
			errCause := _errors.Cause(err)
			switch errCause {
			// http status 400
			case kurima.ErrDuplicatedUser:
				return c.JSON(http.StatusBadRequest, errCause)
			case kurima.ErrBindStruct:
				return c.JSON(http.StatusBadRequest, errCause)
			case kurima.ErrValidateStruct:
				return c.JSON(http.StatusBadRequest, errCause)
			case kurima.ErrorWrongPassword:
				// http status 401
				return c.JSON(http.StatusUnauthorized, kurima.ErrorWrongPassword)
			case kurima.ErrInValidUser:
				return c.JSON(http.StatusUnauthorized, kurima.ErrorWrongPassword)
			case kurima.ErrNotFound:
				return c.JSON(http.StatusNotFound, kurima.ErrNotFound)
			default:
				// http status 500
				return c.JSON(http.StatusInternalServerError, errCause)
			}
		}

		token, err := kurima.GetTokenOnContext(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		c.Response().Header().Set(echo.HeaderAuthorization, token)
		return handler(c)
	}
}

// NewMiddleware .
func NewMiddleware(jwtHash kurima.JWTHash, userRepo kurima.UserRepository) Middleware {
	return Middleware{
		jwtHash:  jwtHash,
		userRepo: userRepo,
	}
}
