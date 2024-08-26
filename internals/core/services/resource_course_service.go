package services

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type ResourceCourseService struct {
	resourceCourseRepository ports.ResourceCourseRepository
	courseRepository         ports.CourseRepository
}

func NewResourceCourseService(rcRepository ports.ResourceCourseRepository, cRepository ports.CourseRepository) *ResourceCourseService {
	return &ResourceCourseService{
		resourceCourseRepository: rcRepository,
		courseRepository:         cRepository,
	}
}

var _ ports.ResourceCourseService = (*ResourceCourseService)(nil)

func (s *ResourceCourseService) AddResourceToCourse(resourceId, courseId uint) (*domain.Course, error) {

	// TODO: prevent duplicate insertion of resource to course

	// TODO: need to check if the resource and course have relationship existing

	resourceCourse, error := s.resourceCourseRepository.AddResourceToCourse(resourceId, courseId)

	if error != nil {
		log.Error(error)
		return nil, error
	}

	course, error := s.courseRepository.FindOne(resourceCourse.Course_id)

	if error != nil {
		log.Error(error)
		return nil, error
	}

	return course, nil
}

func (s *ResourceCourseService) AsignCourseToResources(resources []uint, courseId uint) (*domain.Course, error) {

	var resourceCourse *domain.ResourceCourse
	var error error
	// loop through the resources slice and call repository method to add resources to course
	for i := range resources {
		resourceCourse, error = s.resourceCourseRepository.AddResourceToCourse(resources[i], courseId)
		if error != nil {
			// if error is not nil then return error and log it
			log.Error(error)
			return nil, error
		}
	}
	// get course id from resourceCourse struct and find it in db with the list of resources
	asignedCourseId := resourceCourse.Course_id

	return s.courseRepository.FindOne(asignedCourseId)
}
