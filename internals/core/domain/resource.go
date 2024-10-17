package domain

type Resource struct {
	Id          uint      `json:"id" gorm:"primaryKey;size:256"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Type        string    `json:"resource_type" validate:"required,oneof=video audio document"`
	Url         string    `json:"url"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Teacher_id  int       `json:"teacher_id" validate:"required"`
	Duration    int       `json:"duration"`
	Image       string    `json:"image"`
	Courses     []*Course `gorm:"many2many:resources_courses;constraint:OnDelete:CASCADE"`
}

func NewResource(id uint, title string, description string, resource_type string, url string, price float64, teacher_id int, duration int, image string) *Resource {

	return &Resource{
		Id:          id,
		Title:       title,
		Description: description,
		Type:        resource_type,
		Url:         url,
		Price:       price,
		Teacher_id:  teacher_id,
		Duration:    duration,
		Image:       image,
	}
}
