package services

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	"github.com/n0o01lh/llp/internals/utils"
)

type ResourceService struct {
	resourceRepository ports.ResourceRepository
	ctx                context.Context
}

var _ ports.ResourceService = (*ResourceService)(nil)

func NewResourceService(ctx context.Context, repository ports.ResourceRepository) *ResourceService {
	return &ResourceService{
		resourceRepository: repository,
		ctx:                ctx,
	}
}

func (service *ResourceService) Create(resource *domain.Resource) (*domain.Resource, error) {

	//upload image to cloudinary
	cloudinary := utils.GetCloudinaryInstance(service.ctx)
	imageUrl, publicId, err := utils.UploadImage(cloudinary, service.ctx, resource.Image)

	if err != nil {
		return nil, err
	}

	resource.Image = imageUrl
	resource.PublicId = publicId

	//dates
	resource.CreatedAt = time.Now().UTC()
	resource.UpdatedAt = time.Now().UTC()

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

func (service *ResourceService) ListAllByTeacherId(teacherId uint) ([]*domain.Resource, error) {
	resourceList, err := service.resourceRepository.ListAllByTeacherId(teacherId)

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

	currentResource, _ := service.resourceRepository.FindOne(id)
	cloudinary := utils.GetCloudinaryInstance(service.ctx)

	//check if is neccessary to update image
	if strings.Contains(currentResource.Image, "res.cloudinary.com") &&
		strings.Contains(resource.Image, "base64") {
		//remove previous image
		err := utils.RemoveImage(cloudinary, service.ctx, currentResource.PublicId)

		if err != nil {
			log.Error("Unable to remove cdn image")
			return nil, err
		}
	}

	if strings.Contains(resource.Image, "base64") {
		//upload image to cloudinary
		imageUrl, publicId, err := utils.UploadImage(cloudinary, service.ctx, resource.Image)

		if err != nil {
			return nil, err
		}
		resource.Image = imageUrl
		resource.PublicId = publicId
	}

	//date
	resource.UpdatedAt = time.Now().UTC()

	resourceUpdated, err := service.resourceRepository.Update(id, resource)

	if err != nil {
		return nil, err
	}

	return resourceUpdated, nil
}

func (service *ResourceService) Delete(id uint) error {
	cloudinary := utils.GetCloudinaryInstance(service.ctx)
	currentResource, err := service.resourceRepository.FindOne(id)

	if err != nil {
		return err
	}

	err = service.resourceRepository.Delete(id)

	if err != nil {
		return err
	}

	err = utils.RemoveImage(cloudinary, service.ctx, currentResource.PublicId)

	if err != nil {
		log.Error("Unable to remove user image from cdn")
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

func (service *ResourceService) SalesHistoryByTeacher(teacherId uint) ([]*domain.ResourceSalesHisotry, error) {

	salesHistory, err := service.resourceRepository.SalesHistoryByTeacher(teacherId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return salesHistory, nil
}
