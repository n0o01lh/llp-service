package services

import (
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type CourseService struct {
	courseRepository ports.CourseRepository
}

func NewCourseService(courseRepo ports.CourseRepository) *CourseService {
	return &CourseService{
		courseRepository: courseRepo,
	}
}

var _ ports.CourseService = (*CourseService)(nil)

func (service *CourseService) Create(course *domain.Course) (*domain.Course, error) {
	courseCreated, err := service.courseRepository.Create(course)

	if err != nil {
		return nil, err
	}
	return courseCreated, nil
}

func (service *CourseService) ListAll() ([]*domain.Course, error) {
	courseList, err := service.courseRepository.ListAll()

	if err != nil {
		return nil, err
	}

	return courseList, nil
}

func (service *CourseService) ListAllByTeacherId(teacherId uint) ([]*domain.Course, error) {
	courseList, err := service.courseRepository.ListAllByTeacherId(teacherId)

	if err != nil {
		return nil, err
	}

	return courseList, nil
}

func (service *CourseService) FindOne(id uint) (*domain.Course, error) {
	course, err := service.courseRepository.FindOne(id)

	if err != nil {
		return nil, err
	}

	return course, nil
}

func (service *CourseService) Update(id uint, course *domain.Course) (*domain.Course, error) {
	courseUpdated, err := service.courseRepository.Update(id, course)

	if err != nil {
		return nil, err
	}

	return courseUpdated, nil
}

func (service *CourseService) Delete(id uint) error {
	err := service.courseRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (service *CourseService) SalesHistory(teacherId uint) ([]*domain.CourseSalesHistory, error) {

	salesHistory, err := service.courseRepository.SalesHistory(teacherId)

	if err != nil {
		return nil, err
	}

	return salesHistory, nil
}
