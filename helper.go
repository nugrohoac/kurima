package kurima

import (
	"context"
	"errors"
)

const (
	keyClaim = "claim"
	keyToken = "token"
)

// SetClaimOnContext .
func SetClaimOnContext(ctx context.Context, claim Claim) context.Context {
	return context.WithValue(ctx, keyClaim, claim)
}

// GetClaimOnContext .
func GetClaimOnContext(ctx context.Context) (Claim, error) {
	claim, ok := ctx.Value(keyClaim).(Claim)
	if !ok {
		return Claim{}, errors.New("failed casting key to claim")
	}

	return claim, nil
}

// SetTokenOnContext .
func SetTokenOnContext(ctx context.Context, stringToken string) context.Context {
	return context.WithValue(ctx, keyToken, stringToken)
}

// GetTokenOnContext .
func GetTokenOnContext(ctx context.Context) (string, error) {
	tokenString, ok := ctx.Value(keyToken).(string)
	if !ok {
		return "", errors.New("failed casting key to string")
	}

	return tokenString, nil
}
