package user

import (
	"context"
	"encoding/hex"
	"github.com/nac-project/kurima"
	"hash"

	"gopkg.in/go-playground/validator.v9"



	"github.com/pkg/errors"
)

type service struct {
	userRepo  kurima.UserRepository
	saltStart string
	saltEnd   string
	sha521    hash.Hash
	validator validator.Validate
}

// Register .
func (s service) Register(ctx context.Context, user kurima.User) (kurima.User, error) {
	currentUser, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil && err != kurima.ErrNotFound {
		return kurima.User{}, errors.Wrap(err, "error get user by email")
	}

	if currentUser.Email == user.Email {
		return kurima.User{}, errors.Wrap(kurima.ErrDuplicatedUser, "user already exist")
	}

	user.Password = s.encryptSHA512(user.Password)
	userRegistered, err := s.userRepo.Register(ctx, user)

	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error register user from repository")
	}
	userRegistered.Password = ""

	return userRegistered, nil
}

// Login .
func (s service) Login(ctx context.Context, user kurima.User) (kurima.User, error) {
	user.Password = s.encryptSHA512(user.Password)

	userLogin, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error login")
	}

	if user.Password != userLogin.Password {
		return kurima.User{}, errors.Wrap(kurima.ErrorWrongPassword, "wrong password")
	}

	userLogin.Password = ""

	return userLogin, nil
}

// UpdatePassword .
func (s service) UpdatePassword(ctx context.Context, ID string, user kurima.User) (kurima.User, error) {
	currentUser, err := s.userRepo.GetByID(ctx, ID)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error get user by id")
	}

	if currentUser.Email != user.Email {
		return kurima.User{}, errors.Wrap(kurima.ErrInValidUser, "user is invalid")
	}

	_, err = s.userRepo.UpdatePassword(ctx, ID, user)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error update password")
	}

	currentUser.Password = ""

	return currentUser, nil
}

func (s service) encryptSHA512(str string) string {
	str = s.saltStart + str + s.saltEnd
	s.sha521.Write([]byte(str))
	defer s.sha521.Reset()

	byteHash := s.sha521.Sum(nil)
	stringHash := hex.EncodeToString(byteHash)

	return stringHash
}

// Initiator as a type for constructor
type Initiator func(s *service) *service

// WithUserRepository is used to included repository
func (i Initiator) WithUserRepository(userRepo kurima.UserRepository) Initiator {
	return func(s *service) *service {
		i(s).userRepo = userRepo
		return s
	}
}

// WithSha521 is used to included WithSha521
func (i Initiator) WithSha521(sha521 hash.Hash) Initiator {
	return func(s *service) *service {
		i(s).sha521 = sha521
		return s
	}
}

// WithSaltStart is used to included salt
func (i Initiator) WithSaltStart(saltStart string) Initiator {
	return func(s *service) *service {
		i(s).saltStart = saltStart
		return s
	}
}

// WithSaltEnd is used to included salt
func (i Initiator) WithSaltEnd(saltEnd string) Initiator {
	return func(s *service) *service {
		i(s).saltEnd = saltEnd
		return s
	}
}

// Build is used to build service
func (i Initiator) Build() kurima.UserService {
	return i(&service{})
}

// NewUserService is used to create Initiator
func NewUserService() Initiator {
	return func(s *service) *service {
		return s
	}
}
