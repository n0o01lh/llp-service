package domain

type CourseSalesHistory struct {
	Course_id  uint    `json:"course_id"`
	Title      string  `json:"title"`
	Teacher_id uint    `json:"teacher_id"`
	Amount     float64 `json:"amount"`
}
