package repositories

import (
	//"errors"
	//"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
	db_utils "github.com/n0o01lh/llp/internals/db/db_utils"
	"gorm.io/gorm"
)

type ResourceCourseRepository struct {
	database *gorm.DB
}

func NewResourceCourseRepository(db *gorm.DB) *ResourceCourseRepository {
	return &ResourceCourseRepository{
		database: db,
	}
}

var _ ports.ResourceCourseRepository = (*ResourceCourseRepository)(nil)

func (rc *ResourceCourseRepository) AddResourceToCourse(resourceId, order, courseId uint) (*domain.ResourceCourse, error) {

	recordExists, err := db_utils.IsRecordExists(rc.database,
		"resources_courses", "resource_id = ? and course_id = ?",
		resourceId, courseId)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("recordExists", recordExists)

	var resourceCourse *domain.ResourceCourse = &domain.ResourceCourse{
		ResourceId: resourceId,
		CourseId:   courseId,
		Order:      int(order),
	}

	if recordExists {
		//errorStr := fmt.Sprintf("Resource with id %d already exists in course with id %d", resourceId, courseId)

		//change order
		result := rc.database.Raw(`update resources_courses set "order" = ? where course_id = ? and resource_id = ?`, order, courseId, resourceId).Scan(&resourceCourse)

		if result.Error != nil {
			log.Error("Error updating table resource_course", result.Error)
			return nil, result.Error
		}
		return resourceCourse, nil
		//return nil, errors.New(errorStr)
	} else {

		result := rc.database.Table("resources_courses").Create(&resourceCourse)

		if result.Error != nil {
			log.Error("Error updating table resource_course", result.Error)
			return nil, result.Error
		}

		log.Info("Resource-Course relationship created: ", &resourceCourse)
		return resourceCourse, nil
	}

}

func (rc *ResourceCourseRepository) AsignCourseToResources(resources []uint, courseId uint) (*domain.ResourceCourse, error) {
	return nil, nil
}

func (rc *ResourceCourseRepository) RemoveResourceFromCourse(resourceId, courseId uint) error {

	var resourceCourse *domain.ResourceCourse

	rc.database.Table("resources_courses").Where("resource_id=? and course_id=?", resourceId, courseId).Find(&resourceCourse)

	result := rc.database.Table("resources_courses").Delete(&resourceCourse)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
