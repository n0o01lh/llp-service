package domain

type Course struct {
	Id          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       float64     `json:"price"`
	Resources   []*Resource `gorm:"many2many:resources_courses;"`
}

func NewCourse(id uint, title, description string, price float64) *Course {
	return &Course{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
	}
}
