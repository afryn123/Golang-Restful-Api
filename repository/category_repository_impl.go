package repository

import (
	"afryn123/golang-restful-api/helper"
	"afryn123/golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// Implement saving category to the database
	sql := "INSERT INTO category (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)
	categoryId, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(categoryId)
	return category
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// Implement updating category in the database
	sql := "UPDATE category SET name =? WHERE id =?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	// Implement deleting category from the database
	sql := "DELETE FROM category WHERE id =?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// Implement finding category by ID from the database
	sql := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("Category not found")
	}
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// Implement finding all categories from the database
	sql := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
