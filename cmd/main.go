package main

import (
	"crypto/sha512"
	"database/sql"
	"github.com/nac-project/kurima/internal/_http"
	"github.com/nac-project/kurima/internal/_mysql"
	"github.com/nac-project/kurima/user"
	"log"
	http2 "net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"


	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsnMysql := os.Getenv("DSN")
	mysqlDB, err := sql.Open("mysql", dsnMysql)
	if err != nil {
		logrus.Fatal("FAILED CONNECT TO MYSQL", err.Error())
	}
	if mysqlDB != nil {
		mysqlDB.SetConnMaxLifetime(time.Duration(5) * time.Second)
		mysqlDB.SetMaxIdleConns(3)
		mysqlDB.SetConnMaxLifetime(5)
	}

	userRepository := _mysql.NewUserRepository(mysqlDB)

	userService := user.NewUserService().
		WithUserRepository(userRepository).
		WithSha521(sha512.New()).
		WithSaltStart(os.Getenv("START_SALT")).
		WithSaltEnd(os.Getenv("END_SALT")).
		Build()

	e := echo.New()
	group := e.Group("")

	timeOutString := os.Getenv("TIMEOUT")
	timeOutInt, err := strconv.Atoi(timeOutString)
	if err != nil {
		logrus.Fatal("FAILED LOAD TIMEOUT", err.Error())

	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http2.StatusOK, "pong")
	})

	//_http.NewUserDelivery(e, userService, time.Duration(timeOutInt)*time.Second, *validator.New())

	_http.NewUserDeliveryWithAuth(group, userService, time.Duration(timeOutInt)*time.Second, *validator.New())

	e.Logger.Fatal(e.Start(":3000"))
}