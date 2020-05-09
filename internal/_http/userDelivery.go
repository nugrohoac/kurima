package _http

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/nac-project/kurima"
)

type delivery struct {
	userService kurima.UserService
	timeOut     time.Duration
	_validator  validator.Validate
}

// Register .
func (d delivery) Register(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), d.timeOut)
	defer cancel()

	var user kurima.User
	err := c.Bind(&user)
	if err != nil {
		return errors.Wrap(kurima.ErrBindStruct{Message: err.Error()}, "error binding struct")
	}

	err = d.validate(user)
	if err != nil {
		return errors.Wrap(kurima.ErrValidateStruct{Message: err.Error()}, "error validate struct")
	}

	user, err = d.userService.Register(ctx, user)
	if err != nil {
		return errors.Wrap(err, "error register user")
	}

	return c.JSON(http.StatusCreated, user)
}

// Login .
func (d delivery) Login(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), d.timeOut)
	defer cancel()

	var user kurima.User
	err := c.Bind(&user)
	if err != nil {
		return errors.Wrap(kurima.ErrBindStruct{Message: err.Error()}, "error binding struct")
	}

	user, err = d.userService.Login(ctx, user)
	if err != nil {
		return errors.Wrap(err, "error login")
	}

	return c.JSON(http.StatusOK, user)
}

// UpdatePassword .
func (d delivery) UpdatePassword(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), d.timeOut)
	defer cancel()

	var user kurima.User
	err := c.Bind(&user)
	if err != nil {
		return errors.Wrap(kurima.ErrBindStruct{Message: err.Error()}, "error binding struct")
	}

	ID := c.Param("id")
	user, err = d.userService.UpdatePassword(ctx, ID, user)
	if err != nil {
		return errors.Wrap(err, "error update password user")
	}

	return c.JSON(http.StatusOK, user)
}

func (d delivery) validate(data interface{}) error {
	err := d._validator.Struct(data)
	if err != nil {
		return err
	}

	return nil
}

// NewUserDeliveryWithAuth .
func NewUserDeliveryWithAuth(e *echo.Group, userService kurima.UserService, timeOut time.Duration, _validator validator.Validate) {
	d := delivery{
		userService: userService,
		timeOut:     timeOut,
		_validator:  _validator,
	}

	e.POST("/user/register", d.Register)
	e.POST("/user/login", d.Login)
	e.PUT("/user/update-password/:id", d.UpdatePassword)
}

// NewUserDelivery .
func NewUserDelivery(e *echo.Echo, userService kurima.UserService, timeOut time.Duration, _validator validator.Validate) {
	d := delivery{
		userService: userService,
		timeOut:     timeOut,
		_validator:  _validator,
	}

	e.POST("/user/login", d.Login)
	e.POST("/user/register", d.Register)
}
