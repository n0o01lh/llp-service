package domain

type Course struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewCourse(id uint, title, description string, price float64) *Course {
	return &Course{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       price,
	}
}
