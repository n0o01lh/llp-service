package domain

type Course struct {
	Id          uint        `json:"id"`
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Price       float64     `json:"price"`
	Teacher_id  int         `json:"teacher_id" validate:"required"`
	Resources   []*Resource `gorm:"many2many:resources_courses;"`
}

func NewCourse(id uint, teacher_id int, title, description string, price float64) *Course {
	return &Course{
		Id:          id,
		Title:       title,
		Description: description,
		Teacher_id:  teacher_id,
		Price:       price,
	}
}
