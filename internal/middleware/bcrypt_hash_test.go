package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nac-project/kurima/internal/middleware"
)

func TestNewBcryptHash(t *testing.T) {
	password := "this is example password"

	bcryptHash := middleware.NewBcryptHash().
		WithStartEnd("salt start").
		WithStartEnd("salt end").
		Build()

	passwordHash, err := bcryptHash.Generate(password)
	require.NoError(t, err)
	require.NotEmpty(t, passwordHash)

	err = bcryptHash.Compare(passwordHash, password)
	require.NoError(t, err)
}
