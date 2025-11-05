package model

type Course struct {
	ID        int    `json:"id,omitempty"`         // 自增ID
	CreatedAt string `json:"created_at,omitempty"` // 创建时间
	UpdatedAt string `json:"updated_at,omitempty"` // 更新时间
	CoursePartial
}

type CoursePartial struct {
	UserId          int     `json:"user_id,omitempty"`           // 用户ID
	MedicineName    string  `json:"medicine_name,omitempty"`     // 药物名称
	MedicineImage   *string `json:"medicine_image,omitempty"`    // 药物图片
	MedicineType    int     `json:"medicine_type"`               // 药物方式 0: 内服；1: 外用
	MedicineTiming  int     `json:"medicine_timing,omitempty"`   // 用药时机 0: 不限；1:饭前用药；2: 饭后用药；3: 随餐用药；4: 睡前用药
	CourseStartTime string  `json:"course_start_time,omitempty"` // 用药开始时间
	Status          int     `json:"status"`                      // 方案状态 0: 生效；1: 废弃
}
