package middleware

import (
	"github.com/nac-project/kurima"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHash struct {
	saltStart string
	saltEnd   string
}

// Generate is used to hash password
func (b bcryptHash) Generate(password string) (string, error) {
	password = b.saltStart + password + b.saltEnd
	passwordBytes := []byte(password)
	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

// Compare use to check valid hash with original string
func (b bcryptHash) Compare(hashedPassword, password string) error {
	password = b.saltStart + password + b.saltEnd
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hashedPassword)

	return bcrypt.CompareHashAndPassword(hashedPasswordBytes, passwordBytes)
}

// Initiator as a type for constructor
type Initiator func(s *bcryptHash) *bcryptHash

// WithStartSalt .
func (i Initiator) WithStartSalt(saltStart string) Initiator {
	return func(s *bcryptHash) *bcryptHash {
		i(s).saltStart = saltStart
		return s
	}
}

// WithStartEnd .
func (i Initiator) WithStartEnd(saltEnd string) Initiator {
	return func(s *bcryptHash) *bcryptHash {
		i(s).saltEnd = saltEnd
		return s
	}
}

// Build .
func (i Initiator) Build() kurima.BcryptHash {
	return i(&bcryptHash{})
}

// NewBcryptHash as initialize bcrypt hash
func NewBcryptHash() Initiator {
	return func(s *bcryptHash) *bcryptHash {
		return s
	}
}
