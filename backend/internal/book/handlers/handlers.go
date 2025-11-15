package handlers

import (
	"context"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/domain"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/dto"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/services"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/logger"
)

type BookHandler struct {
	service services.BookService
	log     logger.Logger
}

func NewBookHandler(s services.BookService, l logger.Logger) *BookHandler {
	return &BookHandler{service: s, log: l}
}

func (h *BookHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateBookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	b := &domain.Book{
		OwnerID:  req.OwnerID,
		Title:    req.Title,
		Author:   req.Author,
		Photos:   req.Photos,
		Price:    req.Price,
		Lend:     req.Lend,
		Location: domain.GeoPoint{Lat: req.Lat, Lng: req.Lng},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.service.CreateBook(ctx, b)
	if err != nil {
		h.log.Error("create book failed", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *BookHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.service.GetBook(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(res)
}

func (h *BookHandler) List(c *fiber.Ctx) error {
	limit := int64(20)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := h.service.ListBooks(ctx, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}

	return c.JSON(res)
}

func (h *BookHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing id"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := h.service.DeleteBook(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "delete failed"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
