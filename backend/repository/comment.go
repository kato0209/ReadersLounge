package repository

import (
	"github.com/jmoiron/sqlx"
)

type ICommentRepository interface {
}

type commentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) ICommentRepository {
	return &commentRepository{db}
}
