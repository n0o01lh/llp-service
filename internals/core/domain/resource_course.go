package domain

type ResourceCourse struct {
	Id         uint   `json:"id,omitempty"`
	ResourceId uint   `json:"resource_id"`
	CourseId   uint   `json:"course_id,omitempty"`
	Order      int    `json:"order"`
	Error      string `json:"errors,omitempty" gorm:"-"`
}

type ResourceCourseResponse struct {
	ResourceCourse *ResourceCourse
	Error          string
}
