package domain

import "time"

type ResourceSalesHisotry struct {
	Id          uint      `json:"id"`
	Teacher_id  uint      `json:"teacher_id"`
	Resource_id uint      `json:"resource_id"`
	Date        time.Time `json:"date"`
	Amount      float64   `json:"amount"`
}
