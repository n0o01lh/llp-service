package domain

type Resource struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Type        string    `json:"resource_type"`
	Url         string    `json:"url"`
	Price       float64   `json:"price"`
	Teacher_id  int       `json:"teacher_id"`
	Duration    int       `json:"duration"`
	Image       string    `json:"image"`
	Courses     []*Course `gorm:"many2many:resources_courses;"`
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
