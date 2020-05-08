package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nac-project/kurima"
	"github.com/nac-project/kurima/mocks"
)

func TestMiddleware(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		user := kurima.User{
			Email:  "nugrohoac96@gmail.com",
			Role:   []string{kurima.DefaultRole},
			Status: kurima.StatusInactive,
		}

		userRepoMock := new(mocks.UserRepository)
		userRepoMock.On("GetByEmail", mock.Anything, user.Email).
			Return(user, nil).
			Once()

		jwtHash := NewJWTHash([]byte("secret key"), time.Duration(1)*time.Hour)
		tokenJwt, err := jwtHash.Encode(user)
		assert.NoError(t, err)
		assert.NotEmpty(t, tokenJwt)

		middleware := NewMiddleware(jwtHash, userRepoMock)

		//setup
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+tokenJwt)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		handler := func(c echo.Context) error {
			return nil
		}
		x := middleware.Auth(handler)(ctx)
		fmt.Println(x)
		res, _ := kurima.GetTokenOnContext(ctx.Request().Context())
		fmt.Println(res)
	})
}
