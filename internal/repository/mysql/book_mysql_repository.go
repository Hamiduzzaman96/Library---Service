package mysql

import (
	"context"
	"database/sql"
	"errors"

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

	result, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.ISBN, book.Available) //ExecContext used for INSER,UPDATE and DELETE

	if err != nil {
		return err
	}
	id, _ := result.LastInsertId() //auto increment id return
	book.ID = id

	return nil

}

func (r *BookMySQLRepository) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	query := `SELECT id,title,author,isbn,available FROM books WHERE id = ?`
	book := &domain.Book{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Available,
	)
	if err == sql.ErrNoRows { //database/sql package predefined error,that means query is ok but no rows
		return nil, errors.New("BOOK NOT FOUND") //custom,readable error, 404
	}
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookMySQLRepository) Update(ctx context.Context, book *domain.Book) error {
	query := `UPDATE books SET title = ?, author = ?, available = ? WHERE id =?`

	result, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.Available, book.ID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected() //sql operation -> how many row change
	if rows == 0 {                   //sql query run perfectly but no rows match
		return errors.New("BOOK NOT FOUND")
	}
	return nil

}

func (r *BookMySQLRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM books WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected() //sql operation -> how many row change
	if rows == 0 {
		return errors.New("BOOK NOT FOUND")
	}
	return nil

}
