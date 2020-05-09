package user

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/nac-project/kurima"
)

type service struct {
	validator  validator.Validate
	userRepo   kurima.UserRepository
	bcryptHash kurima.BcryptHash
}

// Register .
func (s service) Register(ctx context.Context, user kurima.User) (kurima.User, error) {
	currentUser, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		switch errors.Cause(err).(type) {
		case kurima.ErrNotFound:
		default:
			return kurima.User{}, errors.Wrap(err, "error get user by email")
		}
	}

	if currentUser.Email == user.Email {
		return kurima.User{}, errors.Wrap(kurima.ErrDuplicated{Message: "user already exist"}, "user already exist")
	}

	user.Password, err = s.bcryptHash.Generate(user.Password)
	if err != nil {
		return kurima.User{}, errors.Wrap(kurima.ErrorAuth{Message: err.Error()}, "error generating password")
	}

	user.Role = append(user.Role, kurima.RoleDefault)
	user.Status = kurima.StatusInactive
	userRegistered, err := s.userRepo.Register(ctx, user)

	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error register user from repository")
	}
	userRegistered.Password = ""

	return userRegistered, nil
}

// Login .
func (s service) Login(ctx context.Context, user kurima.User) (kurima.User, error) {
	userLogin, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error login")
	}

	if err = s.bcryptHash.Compare(userLogin.Password, user.Password); err != nil {
		return kurima.User{}, errors.Wrap(kurima.ErrorAuth{Message: err.Error()}, "error comparing hashed with password")
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
		return kurima.User{}, errors.Wrap(kurima.ErrInValid{Message: "email is different"}, "user is invalid")
	}

	_, err = s.userRepo.UpdatePassword(ctx, ID, user)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error update password")
	}

	currentUser.Password = ""

	return currentUser, nil
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

// WithBcryptHash is used to included BcryptHash
func (i Initiator) WithBcryptHash(bcryptHash kurima.BcryptHash) Initiator {
	return func(s *service) *service {
		i(s).bcryptHash = bcryptHash
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
