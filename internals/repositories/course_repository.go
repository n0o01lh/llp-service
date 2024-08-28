package repositories

import (
	"errors"

	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	"gorm.io/gorm"
)

type CourseRepository struct {
	Database *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		Database: db,
	}
}

var _ ports.CourseRepository = (*CourseRepository)(nil)

func (r *CourseRepository) Create(course *domain.Course) (*domain.Course, error) {

	var newCourse *domain.Course

	r.Database.Create(&course)
	r.Database.Find(&newCourse, course.Id)

	if newCourse == nil {
		return nil, errors.New("course not created")
	}

	return newCourse, nil
}

func (r *CourseRepository) ListAll() ([]*domain.Course, error) {

	var courseList []*domain.Course

	r.Database.Preload("Resources").Find(&courseList)

	if courseList == nil {
		return nil, errors.New("courses not found")
	}

	return courseList, nil
}

func (r *CourseRepository) FindOne(id uint) (*domain.Course, error) {

	var course *domain.Course

	row := r.Database.Preload("Resources").Find(&course, id)

	if row.RowsAffected == 0 {
		return nil, errors.New("course not found")
	}

	return course, nil
}

func (r *CourseRepository) Update(id uint, course *domain.Course) (*domain.Course, error) {

	var updatedCourse *domain.Course

	response := r.Database.Where("id = ?", id).Updates(&course)

	r.Database.Find(&updatedCourse, id)

	if response.RowsAffected == 0 {
		return nil, errors.New("course cannot be updated")
	}

	return updatedCourse, nil
}

func (r *CourseRepository) Delete(id uint) error {

	var course *domain.Course

	result := r.Database.Delete(&course, id)

	if result.RowsAffected == 0 {
		return errors.New("course not found")
	}

	return nil
}
