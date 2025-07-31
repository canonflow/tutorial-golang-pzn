package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-mysql/entity"
	"strconv"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepositoryImpl(db *sql.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: db}
}

func (repo *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repo *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	// Define the query
	query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)

	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	// Read the row
	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}
		return comment, nil
	} else {
		return comment, errors.New("ID: " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repo *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	// Define the query
	query := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	// Read the rows
	defer rows.Close()
	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
