package repository

import (
	"backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type IConnectionRepository interface {
	CreateConnection(ctx echo.Context, followerID, followingID int) error
	DeleteConnection(ctx echo.Context, connectionId int) error
	GetFollowingList(ctx echo.Context, userID int, followingList *[]models.Connection) error
	GetFollowerList(ctx echo.Context, userID int, followerList *[]models.Connection) error
}

type connectionRepository struct {
	db *sqlx.DB
}

func NewConnectionRepository(db *sqlx.DB) IConnectionRepository {
	return &connectionRepository{db}
}

func (cr *connectionRepository) CreateConnection(ctx echo.Context, followerID, followingID int) error {
	query := `INSERT INTO connections (follower_id, following_id) VALUES ($1, $2);`
	_, err := cr.db.ExecContext(ctx.Request().Context(), query, followerID, followingID)
	if err != nil {
		return err
	}

	return nil
}

func (cr *connectionRepository) DeleteConnection(ctx echo.Context, connectionId int) error {
	query := `DELETE FROM connections WHERE connection_id = $1;`
	_, err := cr.db.ExecContext(ctx.Request().Context(), query, connectionId)
	if err != nil {
		return err
	}

	return nil
}

func (cr *connectionRepository) GetFollowingList(ctx echo.Context, userID int, followingList *[]models.Connection) error {
	query := `
		SELECT
			c.connection_id,
			c.follower_id,
			c.following_id,
			u.user_id,
		FROM
			connections c
		INNER JOIN
			users u
		ON
			c.following_id = u.user_id
		WHERE
			c.follower_id = $1
	`
	if err := cr.db.SelectContext(ctx.Request().Context(), followingList, query, userID); err != nil {
		return err
	}

	return nil
}

func (cr *connectionRepository) GetFollowerList(ctx echo.Context, userID int, followerList *[]models.Connection) error {
	query := `
		SELECT
			c.connection_id,
			c.follower_id,
			c.following_id,
			u.user_id,
		FROM
			connections c
		INNER JOIN
			users u
		ON
			c.follower_id = u.user_id
		WHERE
			c.following_id = $1
	`
	if err := cr.db.SelectContext(ctx.Request().Context(), followerList, query, userID); err != nil {
		return err
	}

	return nil
}
