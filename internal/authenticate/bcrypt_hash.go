package authenticate

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
func (b bcryptHash) Compare(passwordHashed, password string) error {
	password = b.saltStart + password + b.saltEnd
	passwordBytes := []byte(password)
	passwordHashedBytes := []byte(passwordHashed)

	return bcrypt.CompareHashAndPassword(passwordHashedBytes, passwordBytes)
}

// Initiator as a type for constructor
type Initiator func(s *bcryptHash) *bcryptHash

// WithSaltStart .
func (i Initiator) WithSaltStart(saltStart string) Initiator {
	return func(s *bcryptHash) *bcryptHash {
		i(s).saltStart = saltStart
		return s
	}
}

// WithSaltEnd .
func (i Initiator) WithSaltEnd(saltEnd string) Initiator {
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
