package server

import (
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	//app.Use(logger.New())

	file, err := os.OpenFile("llp-requests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	serverLogsFile, _ := os.OpenFile("llp-server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, serverLogsFile)
	log.SetOutput(iw)

	resourceRoutes := app.Group("/resource")

	resourceRoutes.Post("/create", s.resourceHandlers.Create)
	resourceRoutes.Get("/list", s.resourceHandlers.ListAll)
	resourceRoutes.Get("/find", s.resourceHandlers.FindOne)
	resourceRoutes.Get("/search", s.resourceHandlers.Search)
	resourceRoutes.Get("/sales/:id", s.resourceHandlers.SalesHistory)
	resourceRoutes.Patch("/update/:id", s.resourceHandlers.Update)
	resourceRoutes.Delete("/delete/:id", s.resourceHandlers.Delete)

	courseRoutes := app.Group("/course")

	courseRoutes.Post("/create", s.courseHandlers.Create)
	courseRoutes.Get("/list", s.courseHandlers.ListAll)
	courseRoutes.Get("/find", s.courseHandlers.FindOne)
	courseRoutes.Get("/sales/:teacher_id", s.courseHandlers.SalesHistory)
	courseRoutes.Patch("/update/:id", s.courseHandlers.Update)
	courseRoutes.Delete("/delete/:id", s.courseHandlers.Delete)
	//resource_course adding resource to course
	courseRoutes.Post("/add-one-resource", s.resourceCourseHandlers.AddResourceToCourse)
	courseRoutes.Post("/add-resources", s.resourceCourseHandlers.AsignCourseToResources)
	courseRoutes.Delete("/remove-resource", s.resourceCourseHandlers.RemoveResourceFromCourse)

	app.Listen(":3000")

}
