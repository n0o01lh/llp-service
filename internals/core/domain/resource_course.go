package domain

type ResourceCourse struct {
	Id          uint   `json:"id"`
	Resource_id uint   `json:"resource_id"`
	Course_id   uint   `json:"course_id"`
	Error       string `json:"errors,omitempty" gorm:"-"`
}

type ResourceCourseResponse struct {
	ResourceCourse *ResourceCourse
	Error          string
}
