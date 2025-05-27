package model

type Plan struct {
	ID         int    `json:"id,omitempty"`
	MedicineID int    `json:"medicine_id,omitempty"`
	Amount     int    `json:"amount,omitempty"`
	Type       string `json:"type,omitempty"`
	PlanTime   string `json:"plan_time,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}
