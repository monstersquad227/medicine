package model

type RecordModel struct {
	ID           int     `json:"id,omitempty"`
	UserID       int     `json:"user_id,omitempty"`
	PlanID       int     `json:"plan_id,omitempty"`
	MedicineName string  `json:"medicine_name,omitempty"`
	ActualTime   string  `json:"actual_time,omitempty"`
	Memo         *string `json:"memo,omitempty"`
	IsChecked    int     `json:"is_checked,omitempty"`
	Status       int     `json:"status,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}
