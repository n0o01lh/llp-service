package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/domain"
)

type ResourceCourseService interface {
	AddResourceToCourse(resourceId, courseId uint) (*domain.Course, error)
	AsignCourseToResources(resources []uint, courseId uint) (*domain.Course, error)
}

type ResourceCourseRepository interface {
	AddResourceToCourse(resourceId, courseId uint) (*domain.ResourceCourse, error)
	AsignCourseToResources(resources []uint, courseId uint) (*domain.ResourceCourse, error)
}

type ResourceCourseHandlers interface {
	AddResourceToCourse(ctx *fiber.Ctx) error
	AsignCourseToResources(ctx *fiber.Ctx) error
}
