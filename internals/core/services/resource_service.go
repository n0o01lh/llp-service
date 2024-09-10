package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type ResourceService struct {
	resourceRepository ports.ResourceRepository
}

var _ ports.ResourceService = (*ResourceService)(nil)

func NewResourceService(repository ports.ResourceRepository) *ResourceService {
	return &ResourceService{
		resourceRepository: repository,
	}
}

func (service *ResourceService) Create(resource *domain.Resource) (*domain.Resource, error) {
	resourceCreated, err := service.resourceRepository.Create(resource)

	if err != nil {
		return nil, err
	}
	return resourceCreated, nil
}

func (service *ResourceService) ListAll() ([]*domain.Resource, error) {
	resourceList, err := service.resourceRepository.ListAll()

	if err != nil {
		return nil, err
	}

	return resourceList, nil
}

func (service *ResourceService) FindOne(id uint) (*domain.Resource, error) {
	resource, err := service.resourceRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (service *ResourceService) Update(id uint, resource *domain.Resource) (*domain.Resource, error) {
	resourceUpdated, err := service.resourceRepository.Update(id, resource)

	if err != nil {
		return nil, err
	}

	return resourceUpdated, nil
}

func (service *ResourceService) Delete(id uint) error {
	err := service.resourceRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (service *ResourceService) Search(criteria string) ([]*domain.Resource, error) {

	resources, err := service.resourceRepository.Search(criteria)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Debug(&resources)

	return resources, nil
}

func (service *ResourceService) SalesHistory(resourceId uint) ([]*domain.ResourceSalesHisotry, error) {

	salesHistory, err := service.resourceRepository.SalesHistory(resourceId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return salesHistory, nil
}
