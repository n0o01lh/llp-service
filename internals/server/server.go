package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type Server struct {
	resourceHandlers       ports.ResourceHandlers
	courseHandlers         ports.CourseHandlers
	resourceCourseHandlers ports.ResourceCourseHandlers
	//middlewares here
	//every hanlders will be here
}

func NewServer(rHandlers ports.ResourceHandlers, cHandlers ports.CourseHandlers, rcHandlers ports.ResourceCourseHandlers) *Server {
	return &Server{
		resourceHandlers:       rHandlers,
		courseHandlers:         cHandlers,
		resourceCourseHandlers: rcHandlers,
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
	//resource_course adding resource to course
	courseRoutes.Post("/add-one-resource", s.resourceCourseHandlers.AddResourceToCourse)
	courseRoutes.Post("/add-resources", s.resourceCourseHandlers.AsignCourseToResources)

	app.Listen(":3000")

}
