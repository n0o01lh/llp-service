package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type Server struct {
	resourceHandlers ports.ResourceHandlers
	//middlewares here
	//every hanlders will be here
}

func NewServer(rHandlers ports.ResourceHandlers) *Server {
	return &Server{
		resourceHandlers: rHandlers,
	}
}

func (s *Server) Initialize() {

	app := fiber.New()

	resourceRoutes := app.Group("/resource")

	resourceRoutes.Post("/create", s.resourceHandlers.Create)
	resourceRoutes.Get("/list", s.resourceHandlers.ListAll)
	resourceRoutes.Get("/find", s.resourceHandlers.FindOne)
	resourceRoutes.Patch("/update/:id", s.resourceHandlers.Update)
	resourceRoutes.Delete("/delete/:id", s.resourceHandlers.Delete)

	app.Listen(":3000")

}
