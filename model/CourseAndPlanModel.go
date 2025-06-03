package model

type CourseAndPlan struct {
	CoursePartial
	PlanPartial
	Frequency        int      `json:"frequency,omitempty"`
	CourseStartTimes []string `json:"course_start_times,omitempty"`
}
