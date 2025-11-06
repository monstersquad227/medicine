package model

type CourseAndPlan struct {
	CoursePartial
	PlanPartial
	CourseID         int      `json:"course_id,omitempty"`
	PlanID           int      `json:"plan_id,omitempty"`
	RecordID         int      `json:"record_id,omitempty"`
	IsChecked        int      `json:"is_checked"`
	Frequency        int      `json:"frequency,omitempty"`
	CourseStartTimes []string `json:"course_start_times,omitempty"`
	PlanTimes        string   `json:"plan_times,omitempty"`
	RecordStatus     int      `json:"record_status,omitempty"`
}

type CourseAndPlanAndRecord struct {
	CoursePartial
	PlanPartial
	RecordModel
	Frequency        int      `json:"frequency,omitempty"`
	CourseStartTimes []string `json:"course_start_times,omitempty"`
}
