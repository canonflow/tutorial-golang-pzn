package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-dependency-injection/helper"
	"golang-dependency-injection/model/domain"
)

type CategoryRepositoryImpl struct {
}

// Diganti ke Implementation buat belajar dependecy injection dengan wire
func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "INSERT INTO categories(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, Sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	Sql := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	Sql := "SELECT * FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, Sql, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	Sql := "SELECT * FROM categories"
	rows, err := tx.QueryContext(ctx, Sql)
	helper.PanicIfError(err)

	var categories []domain.Category
	defer rows.Close()
	for rows.Next() {
		category := domain.Category{}
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
