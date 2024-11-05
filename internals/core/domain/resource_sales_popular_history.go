package domain

type ResourcePopularHistory struct {
	Teacher_id  uint    `json:"teacher_id"`
	Resource_id uint    `json:"resource_id"`
	Title       string  `json:"title"`
	SalesCount  float64 `json:"sales_count"`
}
