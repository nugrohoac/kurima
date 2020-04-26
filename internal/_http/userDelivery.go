package _http

import (
	"context"
	"github.com/nac-project/kurima"
	"net/http"
	"time"

	"github.com/labstack/echo"


	"gopkg.in/go-playground/validator.v9"
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
		return kurima.ErrBindStruct
	}

	err = d.validate(user)
	if err != nil {
		return kurima.ErrValidateStruct
	}

	user, err = d.userService.Register(ctx, user)
	if err != nil {
		return err
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
		return kurima.ErrBindStruct
	}

	user, err = d.userService.Login(ctx, user)
	if err != nil {
		return err
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
		return kurima.ErrBindStruct
	}

	ID := c.Param("id")
	user, err = d.userService.UpdatePassword(ctx, ID, user)
	if err != nil {
		return err
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
}
