package mysql

import (
	"context"
	"database/sql"

	"github.com/Hamiduzzaman96/Library---Service/internal/domain"
)

type BookMySQLRepository struct {
	db *sql.DB
}

func NewBookMySQlRepository(db *sql.DB) *BookMySQLRepository {
	return &BookMySQLRepository{db: db}
}

func (r *BookMySQLRepository) Create(ctx context.Context, book *domain.Book) error {
	query := `INSERT INTO books (title,author,isbn,available) VALUES (?,?,?,?)`

	result, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.ISBN, book.Available)

	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	book.ID = id

	return nil

}
