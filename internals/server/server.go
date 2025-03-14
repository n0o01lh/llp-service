package server

import (
	"io"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/n0o01lh/llp/internals/core/ports"
	"github.com/n0o01lh/llp/internals/middlewares"
)

type Server struct {
	resourceHandlers       ports.ResourceHandlers
	courseHandlers         ports.CourseHandlers
	resourceCourseHandlers ports.ResourceCourseHandlers
	userHandlers           ports.UserHandlers
	//middlewares here
	//every hanlders will be here
}

func NewServer(rHandlers ports.ResourceHandlers, cHandlers ports.CourseHandlers, rcHandlers ports.ResourceCourseHandlers, uHandlers ports.UserHandlers) *Server {
	return &Server{
		resourceHandlers:       rHandlers,
		courseHandlers:         cHandlers,
		resourceCourseHandlers: rcHandlers,
		userHandlers:           uHandlers,
	}
}

func (s *Server) Initialize() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

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

	resourceRoutes := app.Group("/resource", middlewares.AuthenticationMiddleware)

	resourceRoutes.Post("/create", middlewares.AuthorizationTeacherMiddleware, s.resourceHandlers.Create)
	resourceRoutes.Get("/list", s.resourceHandlers.ListAll)
	resourceRoutes.Get("/list-by-teacher", s.resourceHandlers.ListAllByTeacherId)
	resourceRoutes.Get("/find", s.resourceHandlers.FindOne)
	resourceRoutes.Get("/search", s.resourceHandlers.Search)
	resourceRoutes.Get("/sales/:id", middlewares.AuthorizationTeacherMiddleware, s.resourceHandlers.SalesHistory)
	resourceRoutes.Get("/sales-by-teacher", middlewares.AuthorizationTeacherMiddleware, s.resourceHandlers.SalesHistoryByTeacher)
	resourceRoutes.Get("/sales-count-by-teacher", middlewares.AuthorizationTeacherMiddleware, s.resourceHandlers.SalesCountHistoryByTeacher)
	resourceRoutes.Patch("/update/:id", middlewares.AuthorizationTeacherMiddleware, s.resourceHandlers.Update)
	resourceRoutes.Delete("/delete/:id", middlewares.AuthorizationTeacherMiddleware, s.resourceHandlers.Delete)

	courseRoutes := app.Group("/course", middlewares.AuthenticationMiddleware)

	courseRoutes.Post("/create", middlewares.AuthorizationTeacherMiddleware, s.courseHandlers.Create)
	courseRoutes.Get("/list", s.courseHandlers.ListAll)
	courseRoutes.Get("/list-by-teacher", s.courseHandlers.ListAllByTeacherId)
	courseRoutes.Get("/find", s.courseHandlers.FindOne)
	courseRoutes.Get("/sales", middlewares.AuthorizationTeacherMiddleware, s.courseHandlers.SalesHistory)
	courseRoutes.Patch("/update/:id", middlewares.AuthorizationTeacherMiddleware, s.courseHandlers.Update)
	courseRoutes.Delete("/delete/:id", middlewares.AuthorizationTeacherMiddleware, s.courseHandlers.Delete)
	//resource_course adding resource to course
	courseRoutes.Post("/add-one-resource", middlewares.AuthorizationTeacherMiddleware, s.resourceCourseHandlers.AddResourceToCourse)
	courseRoutes.Post("/add-resources", middlewares.AuthorizationTeacherMiddleware, s.resourceCourseHandlers.AsignCourseToResources)
	courseRoutes.Delete("/remove-resource", middlewares.AuthorizationTeacherMiddleware, s.resourceCourseHandlers.RemoveResourceFromCourse)
	courseRoutes.Get("/search", s.courseHandlers.Search)

	userRoutes := app.Group("/user")

	userRoutes.Post("/register", s.userHandlers.Register)
	userRoutes.Post("/login", s.userHandlers.Login)

	app.Listen(":3000")

}
