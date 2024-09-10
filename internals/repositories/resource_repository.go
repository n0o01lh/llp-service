package repositories

import (
	"errors"

	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	"gorm.io/gorm"
)

type ResourceRepository struct {
	Database *gorm.DB
}

func NewResourceRepository(database *gorm.DB) *ResourceRepository {
	return &ResourceRepository{
		Database: database,
	}
}

var _ ports.ResourceRepository = (*ResourceRepository)(nil)

func (r *ResourceRepository) Create(resource *domain.Resource) (*domain.Resource, error) {

	var newResource *domain.Resource

	r.Database.Create(&resource)
	r.Database.Find(&newResource, resource.Id)

	if newResource == nil {
		return nil, errors.New("resource not created")
	}

	return newResource, nil
}

func (r *ResourceRepository) ListAll() ([]*domain.Resource, error) {

	var resourceList []*domain.Resource

	r.Database.Preload("Courses").Find(&resourceList)

	if resourceList == nil {
		return nil, errors.New("resources not found")
	}

	return resourceList, nil
}

func (r *ResourceRepository) FindOne(id uint) (*domain.Resource, error) {

	var resource *domain.Resource

	row := r.Database.Preload("Courses").Find(&resource, id)

	if row.RowsAffected == 0 {
		return nil, errors.New("resource not found")
	}

	return resource, nil
}

func (r *ResourceRepository) Update(id uint, resource *domain.Resource) (*domain.Resource, error) {

	var updatedResource *domain.Resource

	r.Database.Where("id = ?", id).Updates(&resource)

	r.Database.Find(&updatedResource, id)

	if updatedResource == nil {
		return nil, errors.New("resource cannot be updated")
	}

	return updatedResource, nil
}

func (r *ResourceRepository) Delete(id uint) error {

	var resource *domain.Resource

	result := r.Database.Delete(&resource, id)

	if result.RowsAffected == 0 {
		return errors.New("resource not found")
	}

	return nil
}

func (r *ResourceRepository) Search(criteria string) ([]*domain.Resource, error) {

	var resources []*domain.Resource

	result := r.Database.Where("lower(title) LIKE lower(?)", "%"+criteria+"%").Table("resources").Find(&resources)

	if result.Error != nil {
		return nil, errors.New("error performing search on database")
	}

	return resources, nil
}

func (r *ResourceRepository) SalesHistory(resourceId uint) ([]*domain.ResourceSalesHisotry, error) {

	var salesHistory []*domain.ResourceSalesHisotry

	result := r.Database.Table("sales_history").Where("resource_id = ?", resourceId).Find(&salesHistory)

	if result.Error != nil {
		return nil, errors.New("error performing search on database")
	}

	return salesHistory, nil
}
