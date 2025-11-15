package server

import (
	"context"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/handlers"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/repository"
	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/services"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/config"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	app *fiber.App
	cfg *config.Config
	log logger.Logger
}

func New(cfg *config.Config, log logger.Logger) (*Server, error) {
	app := fiber.New()

	// create mongodb client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URI))
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.Mongo.Database)

	repo := repository.NewMongoRepo(db, log)
	service := services.NewBookService(repo, log)
	h := handlers.NewBookHandler(service, log)

	app.Post("/books", h.Create)
	app.Get("/books", h.List)
	app.Get("/books/:id", h.Get)
	app.Delete("/books/:id", h.Delete)
	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })

	return &Server{app: app, cfg: cfg, log: log}, nil
}

func (s *Server) Run() {
	s.log.Info("Book service running on " + s.cfg.Server.Addr)
	s.app.Listen(s.cfg.Server.Addr)
}
