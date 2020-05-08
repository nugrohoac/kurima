package authenticate_test

import (
	"testing"
	"time"

	"github.com/nac-project/kurima/internal/authenticate"

	"github.com/stretchr/testify/require"

	"github.com/nac-project/kurima"
)

func TestNewJWTHash(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		secretKey := []byte("this is secret key")
		duration := time.Duration(1) * time.Hour

		jwtHash := authenticate.NewJWTHash(secretKey, duration)
		user := kurima.User{
			Email:  "nugrohoac96@gmail.com",
			Role:   []string{kurima.DefaultRole},
			Status: kurima.StatusInactive,
		}

		tokenJwt, err := jwtHash.Encode(user)
		require.NoError(t, err)
		require.NotEmpty(t, tokenJwt)

		var c kurima.Claim
		valid, err := jwtHash.Decode(tokenJwt, &c)
		require.NoError(t, err)
		require.True(t, valid)
	})
}
