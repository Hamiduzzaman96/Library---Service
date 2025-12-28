package repository

import (
	"context"

	"github.com/Hamiduzzaman96/Library---Service/internal/domain"
)

type BookRepository interface {
	Create(ctx context.Context, book *domain.Book) error         //single return
	GetByID(ctx context.Context, id int64) (*domain.Book, error) //multiple return
	Update(ctx context.Context, book *domain.Book) error
	Delete(ctx context.Context, id int64) error
}
