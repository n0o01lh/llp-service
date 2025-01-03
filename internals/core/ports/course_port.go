package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/domain"
)

type CourseService interface {
	Create(course *domain.Course) (*domain.Course, error)
	ListAll() ([]*domain.Course, error)
	ListAllByTeacherId(teacherId uint, pagination *domain.Pagination) (*domain.Pagination, error)
	FindOne(id uint) (*domain.Course, error)
	Update(id uint, course *domain.Course) (*domain.Course, error)
	Delete(id uint) error
	SalesHistory(teacherId uint) ([]*domain.CourseSalesHistory, error)
	Search(criteria string, teacherId uint, pagination *domain.Pagination) (*domain.Pagination, error)
}

type CourseRepository interface {
	Create(course *domain.Course) (*domain.Course, error)
	ListAll() ([]*domain.Course, error)
	ListAllByTeacherId(teacherId uint, pagination *domain.Pagination) (*domain.Pagination, error)
	FindOne(id uint) (*domain.Course, error)
	Update(id uint, course *domain.Course) (*domain.Course, error)
	Delete(id uint) error
	SalesHistory(teacherId uint) ([]*domain.CourseSalesHistory, error)
	Search(criteria string, teacherId uint, pagination *domain.Pagination) (*domain.Pagination, error)
}

type CourseHandlers interface {
	Create(context *fiber.Ctx) error
	ListAll(context *fiber.Ctx) error
	ListAllByTeacherId(context *fiber.Ctx) error
	FindOne(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	SalesHistory(context *fiber.Ctx) error
	Search(context *fiber.Ctx) error
}
