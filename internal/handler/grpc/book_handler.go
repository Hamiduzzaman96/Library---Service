package grpc

import (
	"context"

	pb "github.com/Hamiduzzaman96/Library---Service/Library---Service/proto/bookpb"
	"github.com/Hamiduzzaman96/Library---Service/internal/domain"
	"github.com/Hamiduzzaman96/Library---Service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "go.opentelemetry.io/otel/codes"
	// "google.golang.org/genproto/googleapis/rpc/status"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

type BookHandler struct {
	usecase *usecase.BookUsecase
	pb.UnimplementedBookServiceServer
}

func NewBookHandler(u *usecase.BookUsecase) *BookHandler {
	return &BookHandler{usecase: u}
}

func (h *BookHandler) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.Empty, error) {
	book := &domain.Book{
		Title:     req.Title,
		Author:    req.Author,
		ISBN:      req.Isbn,
		Available: true,
	}

	err := h.usecase.Create(ctx, book)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *BookHandler) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.BookResponse, error) {
	book, err := h.usecase.GetByID(ctx, req.Id)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.BookResponse{
		Id:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Isbn:      book.ISBN,
		Available: book.Available,
	}, nil
}

func (h *BookHandler) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.Empty, error) {
	book := &domain.Book{
		ID:        req.Id,
		Title:     req.Title,
		Author:    req.Author,
		Available: req.Available,
	}

	err := h.usecase.Update(ctx, book)

	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	return &pb.Empty{}, nil
}

func (h *BookHandler) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.Empty, error) {
	err := h.usecase.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.Empty{}, nil
}
