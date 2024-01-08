package suggestion

import (
	"blinders/packages/suggestion"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Service struct {
	suggester suggestion.Suggester
	app       *fiber.App
	*ServiceConfig
}

func NewTransporter(suggester suggestion.Suggester, cfg *ServiceConfig) (*Service, error) {
	fiberCfg := fiber.Config{
		ReadTimeout:     time.Second * 5,
		WriteTimeout:    time.Second * 5,
		ReadBufferSize:  2e5,
		WriteBufferSize: 2e5,
	}
	fiberApp := fiber.New(fiberCfg)

	s := &Service{
		suggester:     suggester,
		app:           fiberApp,
		ServiceConfig: cfg,
	}
	defer s.initRoute()
	return s, nil
}

func (s *Service) initRoute() {
	s.app.Use(cors.New())
	s.app.Post("/api/suggest/text", s.HandleTextSuggestion())
	s.app.Post("/api/suggest/chat", s.HandleChatSuggestion())
	s.app.Get("/ping", func(c *fiber.Ctx) error {
		_, err := c.WriteString("pong")
		return err
	})
}

func (s *Service) Listen() error {
	ln, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		return err
	}
	return s.app.Listener(ln)
}
