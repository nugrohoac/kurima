package _mysql

import (
	"context"
	"database/sql"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/nac-project/kurima"
)

type repository struct {
	db *sql.DB
}

// GetByEmail .
func (r repository) GetByEmail(ctx context.Context, email string) (kurima.User, error) {
	query, args, err := sq.Select(
		"id",
		"email",
		"password",
		"role",
		"status",
	).From("user").
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error parsing builder into sql")
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var (
		user  kurima.User
		roles string
	)
	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&roles,
		&user.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return kurima.User{}, errors.Wrap(kurima.ErrNotFound{Message: err.Error()}, "user is not found")
		}

		return kurima.User{}, errors.Wrap(err, "error scan user")
	}

	user.Role = strings.Split(roles, ",")

	return user, nil
}

// Register .
func (r repository) Register(ctx context.Context, user kurima.User) (kurima.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error create transaction db")
	}

	roles := strings.Join(user.Role, ",")
	timeNow := time.Now().UTC()
	user.ID = uuid.NewV4().String()

	query, args, err := sq.Insert("user").
		Columns(
			"id",
			"email",
			"password",
			"role",
			"status",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		).Values(
		user.ID,
		user.Email,
		user.Password,
		roles,
		user.Status,
		timeNow,
		timeNow,
		kurima.RoleAdmin,
		kurima.RoleAdmin,
	).ToSql()

	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error convert builder to sql")
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error execute query sql")
	}

	if errCommit := tx.Commit(); errCommit != nil {
		logrus.Error("Error commit : ", errCommit)

		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error("Error rollback at commit")
		}

		return kurima.User{}, errors.Wrap(errCommit, "error commit")
	}

	return user, nil
}

// Login .
func (r repository) Login(ctx context.Context, user kurima.User) (kurima.User, error) {
	query, args, err := sq.Select(
		"id",
		"email",
		"password",
		"role",
		"status",
	).From("user").
		Where(sq.Eq{"email": user.Email}).
		Where(sq.Eq{"password": user.Password}).
		ToSql()
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error parsing builder into sql")
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var roles string

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&roles,
		&user.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return kurima.User{}, errors.Wrap(kurima.ErrNotFound{Message: "user is not found"}, "user is not found")
		}

		return kurima.User{}, errors.Wrap(err, "error scan user")
	}

	user.Role = strings.Split(roles, ",")

	return user, nil
}

// UpdatePassword .
func (r repository) UpdatePassword(ctx context.Context, ID string, user kurima.User) (kurima.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error create transaction db")
	}

	timeNow := time.Now().UTC()
	user.Status = kurima.StatusActive

	query, args, err := sq.Update("user").
		SetMap(sq.Eq{
			"password":   user.Password,
			"status":     user.Status,
			"updated_by": user.Email,
			"updated_at": timeNow,
		}).
		Where(sq.Eq{"id": ID}).
		ToSql()

	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error convert builder to sql")
	}

	result, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error execute query sql")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return kurima.User{}, errors.Wrap(kurima.ErrNoRowAffected{Message: err.Error()}, "error check row affected")
	}

	if rowsAffected == 0 {
		return kurima.User{}, errors.Wrap(kurima.ErrNoRowAffected{Message: err.Error()}, "error row affected is zero")
	}

	if errCommit := tx.Commit(); errCommit != nil {
		logrus.Error("Error commit : ", errCommit)

		if errRollback := tx.Rollback(); errRollback != nil {
			logrus.Error("Error rollback at commit")
		}

		return kurima.User{}, errors.Wrap(errCommit, "error commit")
	}

	return user, nil
}

// GetByID .
func (r repository) GetByID(ctx context.Context, ID string) (kurima.User, error) {
	query, args, err := sq.Select(
		"id",
		"email",
		"password",
		"role",
		"status",
	).From("user").
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return kurima.User{}, errors.Wrap(err, "error parsing builder into sql")
	}

	row := r.db.QueryRowContext(ctx, query, args...)

	var (
		user  kurima.User
		roles string
	)
	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&roles,
		&user.Status,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return kurima.User{}, errors.Wrap(kurima.ErrNotFound{Message: "user is not found"}, "user is not found")
		}

		return kurima.User{}, errors.Wrap(err, "error scan user")
	}

	user.Role = strings.Split(roles, ",")

	return user, nil
}

// NewUserRepository .
func NewUserRepository(db *sql.DB) kurima.UserRepository {
	return repository{
		db: db,
	}
}
