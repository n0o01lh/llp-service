package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0o01lh/llp/internals/core/domain"
)

type ResourceService interface {
	Create(resource *domain.Resource) (*domain.Resource, error)
	ListAll() ([]*domain.Resource, error)
	ListAllByTeacherId(teacherId uint) ([]*domain.Resource, error)
	FindOne(id uint) (*domain.Resource, error)
	Update(id uint, resource *domain.Resource) (*domain.Resource, error)
	Delete(id uint) error
	Search(criteria string) ([]*domain.Resource, error)
	SalesHistory(resourceId uint) ([]*domain.ResourceSalesHisotry, error)
	SalesHistoryByTeacher(teacherId uint) ([]*domain.ResourceSalesHisotry, error)
}

type ResourceRepository interface {
	Create(resource *domain.Resource) (*domain.Resource, error)
	ListAll() ([]*domain.Resource, error)
	ListAllByTeacherId(teacherId uint) ([]*domain.Resource, error)
	FindOne(id uint) (*domain.Resource, error)
	Update(id uint, resource *domain.Resource) (*domain.Resource, error)
	Delete(id uint) error
	Search(criteria string) ([]*domain.Resource, error)
	SalesHistory(resourceId uint) ([]*domain.ResourceSalesHisotry, error)
	SalesHistoryByTeacher(teacherId uint) ([]*domain.ResourceSalesHisotry, error)
}

type ResourceHandlers interface {
	Create(context *fiber.Ctx) error
	ListAll(context *fiber.Ctx) error
	ListAllByTeacherId(context *fiber.Ctx) error
	FindOne(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
	Search(context *fiber.Ctx) error
	SalesHistory(context *fiber.Ctx) error
	SalesHistoryByTeacher(context *fiber.Ctx) error
}
