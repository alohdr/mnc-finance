package internal

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"mnc-finance/config"
	"mnc-finance/internal/handler"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *fiber.App
	config     *config.Config
}

func InitServer(cfg *config.Config) *Server {

	f := fiber.New()

	f.Use(cors.New(), logger.New(), recover.New())
	s := new(Server)
	s.httpServer = f
	s.config = cfg
	return s
}

func (s *Server) handleShutdown(server *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return err
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr: fmt.Sprintf("%v:%v", s.config.Host, s.config.Port),
	}

	handler := handler.NewHandler(RegisterUc(s.config))
	api := s.httpServer
	handler.Mount(api)

	go func() {
		log.Printf("Server starting on port :%d\n", s.config.Port)
		if err := s.httpServer.Listen(fmt.Sprintf(":%d", s.config.Port)); err != nil {
			log.Panic(err)
		}
	}()

	return s.handleShutdown(server)

}
