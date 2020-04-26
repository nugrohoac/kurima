package kurima

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dgrijalva/jwt-go"
)

func TestSetGetClaimOnContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		expiredTime := time.Now().Add(1 * time.Hour)
		claim := Claim{
			User: User{
				Email:  "nugrohoac96@gmail.com",
				Role:   []string{DefaultRole},
				Status: StatusInactive,
			},
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiredTime.Unix(),
			},
		}

		ctx = SetClaimOnContext(ctx, claim)

		claimOnContext, err := GetClaimOnContext(ctx)
		assert.NoError(t, err)
		assert.Equal(t, claim, claimOnContext)
	})
}

func TestSetGetTokenOnContext(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		token := "token jwt"
		ctx = SetTokenOnContext(ctx, token)

		tokenString, err := GetTokenOnContext(ctx)
		assert.NoError(t, err)
		assert.Equal(t, token, tokenString)
	})
}
