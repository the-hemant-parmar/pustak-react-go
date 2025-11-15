package main

import (
	"log"

	"github.com/the-hemant-parmar/pustak-react-go/backend/internal/book/server"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/config"
	"github.com/the-hemant-parmar/pustak-react-go/backend/pkg/common/logger"
)

func main() {
	cfg := config.Load("config/book.yaml")
	logg := logger.New(cfg)

	srv, err := server.New(cfg, logg)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	srv.Run()
}
