package services

import (
	"context"

	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/domain"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/repository"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/logger"
)

type BookService interface {
	CreateBook(ctx context.Context, b *domain.Book) (*domain.Book, error)
	GetBook(ctx context.Context, id string) (*domain.Book, error)
	ListBooks(ctx context.Context, limit int64) ([]*domain.Book, error)
	DeleteBook(ctx context.Context, id string) error
}

type bookService struct {
	repo repository.BookRepository
	log  logger.Logger
}

func NewBookService(r repository.BookRepository, l logger.Logger) BookService {
	return &bookService{repo: r, log: l}
}

func (s *bookService) CreateBook(ctx context.Context, b *domain.Book) (*domain.Book, error) {
	// basic business rules can go here
	bCopy := *b
	return s.repo.Create(ctx, &bCopy)
}

func (s *bookService) GetBook(ctx context.Context, id string) (*domain.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *bookService) ListBooks(ctx context.Context, limit int64) ([]*domain.Book, error) {
	return s.repo.List(ctx, limit)
}

func (s *bookService) DeleteBook(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
