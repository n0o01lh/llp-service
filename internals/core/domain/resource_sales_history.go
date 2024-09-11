package domain

type ResourceSalesHisotry struct {
	Teacher_id  uint    `json:"teacher_id"`
	Resource_id uint    `json:"resource_id"`
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
}
