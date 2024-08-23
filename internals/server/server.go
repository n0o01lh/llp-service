package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type Server struct {
	resourceHandlers ports.ResourceHandlers
	courseHandlers   ports.CourseHandlers
	//middlewares here
	//every hanlders will be here
}

func NewServer(rHandlers ports.ResourceHandlers, cHandlers ports.CourseHandlers) *Server {
	return &Server{
		resourceHandlers: rHandlers,
		courseHandlers:   cHandlers,
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

	courseRoutes := app.Group("/course")

	courseRoutes.Post("/create", s.courseHandlers.Create)
	courseRoutes.Get("/list", s.courseHandlers.ListAll)
	courseRoutes.Get("/find", s.courseHandlers.FindOne)
	courseRoutes.Patch("/update/:id", s.courseHandlers.Update)
	courseRoutes.Delete("/delete/:id", s.courseHandlers.Delete)

	app.Listen(":3000")

}
