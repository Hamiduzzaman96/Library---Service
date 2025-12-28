package usecase

import (
	"context"

	"github.com/Hamiduzzaman96/Library---Service/internal/domain"
	"github.com/Hamiduzzaman96/Library---Service/internal/repository"
)

type BookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(r repository.BookRepository) *BookUsecase {
	return &BookUsecase{repo: r}
}

func (u *BookUsecase) Create(ctx context.Context, book *domain.Book) error {
	return u.repo.Create(ctx, book)
}

func (u *BookUsecase) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *BookUsecase) Update(ctx context.Context, book *domain.Book) error {
	return u.repo.Update(ctx, book)
}

func (u *BookUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
