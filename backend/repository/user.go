package repository

import (
	"backend/db"
	"backend/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IUserRepository interface {
	CreateUser(ctx echo.Context, user *models.User) error
	GetUserByIdentifier(ctx echo.Context, user *models.User, identifier string) error
	GetUserByUserID(ctx echo.Context, user *models.User, userID int) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(ctx echo.Context, user *models.User) error {
	c := ctx.Request().Context()
	var userID int
	if err := db.Tx(c, ur.db, func(tx *sqlx.Tx) error {

		if err := tx.QueryRowContext(c, "INSERT INTO users DEFAULT VALUES RETURNING user_id;").Scan(&userID); err != nil {
			return errors.WithStack(err)
		}
		user.UserID = userID

		var sqlString string
		var sqlArgs []interface{}
		if user.ProfileImage == "" {
			sqlString = "INSERT INTO user_details ( user_id, name, profile_text ) VALUES ($1, $2, $3) RETURNING profile_image;"
			sqlArgs = append(sqlArgs, userID, user.Name, user.ProfileText)
		} else {
			sqlString = "INSERT INTO user_details ( user_id, name, profile_text, profile_image ) VALUES ($1, $2, $3, $4) RETURNING profile_image;"
			sqlArgs = append(sqlArgs, userID, user.Name, user.ProfileText, user.ProfileImage)
		}

		var profileImage string
		err := tx.QueryRowContext(
			c,
			sqlString,
			sqlArgs...,
		).Scan(&profileImage)
		if err != nil {
			return errors.WithStack(err)
		}
		user.ProfileImage = profileImage

		_, err = tx.ExecContext(
			c,
			`
			INSERT INTO user_auths (user_id, identity_type, identifier, credential)
			VALUES ($1, $2, $3, $4);
		`,
			userID,
			user.IdentityType,
			user.Identifier,
			user.Credential,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil

	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (ur *userRepository) GetUserByIdentifier(ctx echo.Context, user *models.User, identifier string) error {
	c := ctx.Request().Context()
	if err := ur.db.GetContext(
		c,
		user,
		`
		select
			users.user_id,
			ua.identifier,
			ua.credential,
			ud.name,
			ud.profile_image
		from users
		inner join user_auths ua using (user_id)
		inner join user_details ud using (user_id)
		where ua.identifier = $1 ;
		`,
		identifier,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.WithStack(errors.New("user is not found"))
		}
		return errors.WithStack(err)
	}
	return nil
}

func (ur *userRepository) GetUserByUserID(ctx echo.Context, user *models.User, userID int) error {
	c := ctx.Request().Context()
	if err := ur.db.GetContext(
		c,
		user,
		`
		select
			users.user_id,
			ud.name,
			ud.profile_image
		from users
		inner join user_details ud using (user_id)
		where users.user_id = $1 ;
		`,
		userID,
	); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
