package services

import (
	"fmt"
	"sync"

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

func (s *ResourceCourseService) AddResourceToCourse(resourceId, courseId uint) (*domain.ResourceCourse, error) {

	result, error := s.resourceCourseRepository.AddResourceToCourse(resourceId, courseId)

	if error != nil {
		log.Error(error)
		return nil, error
	}

	return result, nil
}

func (s *ResourceCourseService) AsignCourseToResources(resources []any, courseId uint) ([]*domain.ResourceCourseResponse, error) {

	var waitGroup sync.WaitGroup
	/* 	var resourceCourse *domain.ResourceCourse
	   	var err error */
	resourceCourses := make([]*domain.ResourceCourseResponse, len(resources))
	errorJoin := make(chan error, len(resources))

	waitGroup.Add(len(resources))

	// loop through the resources slice and call repository method to add resources to course
	for i := range resources {

		go func(index int) {

			resourceCourse, err := s.resourceCourseRepository.AddResourceToCourse(uint(resources[index].(float64)), courseId)

			resourceCourses[index] = &domain.ResourceCourseResponse{ResourceCourse: resourceCourse, Error: fmt.Sprintf("%v", err)}

			if err != nil {
				log.Error(err)
				errorJoin <- err
			}
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()

	/* 	if resourceCourse == nil {
		return nil, errors.New("All resources id are part of the course")
	} */

	// get course id from resourceCourse struct and find it in db with the list of resources
	//asignedCourseId := resourceCourse.Course_id

	/* 	//TODO: how to return a list of resoucrces duplicated in response when other resources are added
	   	if errorJoin != nil {
	   		// if error is not nil then return error and log it
	   		resourceCourse.Errors = (<-errorJoin).Error()
	   	} */

	return resourceCourses, nil
}

func (s *ResourceCourseService) RemoveResourceFromCourse(resourceId, courseId uint) error {

	course, error := s.courseRepository.FindOne(courseId)

	if error != nil {
		return error
	}

	isResourceInCourse := false

	for _, value := range course.Resources {
		if value.Id == resourceId {
			isResourceInCourse = true
			break
		}
	}

	if isResourceInCourse {
		error = s.resourceCourseRepository.RemoveResourceFromCourse(resourceId, courseId)
	}

	if error != nil {
		return error
	}

	return nil
}
