package repository

import (
	"backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type IConnectionRepository interface {
	CreateConnection(ctx echo.Context, followerID, followingID int) (models.Connection, error)
	DeleteConnection(ctx echo.Context, connectionId int) error
	GetFollowingConnections(ctx echo.Context, userID int, followingList *[]models.Connection) error
	GetFollowerConnections(ctx echo.Context, userID int, followerList *[]models.Connection) error
}

type connectionRepository struct {
	db *sqlx.DB
}

func NewConnectionRepository(db *sqlx.DB) IConnectionRepository {
	return &connectionRepository{db}
}

func (cr *connectionRepository) CreateConnection(ctx echo.Context, followerID, followingID int) (models.Connection, error) {
	query := `WITH inserted_connection AS (
				INSERT INTO connections (follower_id, following_id) 
				VALUES ($1, $2) 
				RETURNING connection_id, follower_id
			)
			SELECT 
				ic.connection_id, 
				u.user_id, 
				ud.name, 
				ud.profile_image
			FROM 
				inserted_connection ic
			JOIN 
				user_details ud ON ic.follower_id = ud.user_id
			JOIN 
				users u ON ud.user_id = u.user_id`
	
	var connection models.Connection
	err := cr.db.QueryRowContext(ctx.Request().Context(), query, followerID, followingID).Scan(
		&connection.ConnectionID,
		&connection.Follower.UserID,
		&connection.Follower.Name,
		&connection.Follower.ProfileImage.FileName,
	)
	if err != nil {
		return models.Connection{}, err
	}
	return connection, nil
}

func (cr *connectionRepository) DeleteConnection(ctx echo.Context, connectionId int) error {
	query := `DELETE FROM connections WHERE connection_id = $1;`
	_, err := cr.db.ExecContext(ctx.Request().Context(), query, connectionId)
	if err != nil {
		return err
	}

	return nil
}

func (cr *connectionRepository) GetFollowingConnections(ctx echo.Context, userID int, followingConnections *[]models.Connection) error {
	query := `
		SELECT
			c.connection_id,
			c.following_id,
			ud.name,
			ud.profile_image
		FROM
			connections c
		INNER JOIN
			users u ON u.user_id = c.following_id
		INNER JOIN
			user_details ud ON u.user_id = ud.user_id
		WHERE
			c.follower_id = $1
	`

	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		connection := models.Connection{}
		err := rows.Scan(
			&connection.ConnectionID,
			&connection.Following.UserID,
			&connection.Following.Name,
			&connection.Following.ProfileImage.FileName,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		*followingConnections = append(*followingConnections, connection)
	}
	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (cr *connectionRepository) GetFollowerConnections(ctx echo.Context, userID int, followerConnections *[]models.Connection) error {
	query := `
		SELECT
			c.connection_id,
			c.follower_id,
			ud.name,
			ud.profile_image
		FROM
			connections c
		INNER JOIN
			users u ON u.user_id = c.follower_id
		INNER JOIN
			user_details ud ON u.user_id = ud.user_id
		WHERE
			c.following_id = $1
	`

	rows, err := cr.db.QueryContext(ctx.Request().Context(), query, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	defer rows.Close()

	for rows.Next() {
		connection := models.Connection{}
		err := rows.Scan(
			&connection.ConnectionID,
			&connection.Follower.UserID,
			&connection.Follower.Name,
			&connection.Follower.ProfileImage.FileName,
		)
		if err != nil {
			return errors.WithStack(err)
		}
		*followerConnections = append(*followerConnections, connection)
	}
	if err := rows.Err(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
