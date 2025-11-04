package model

type User struct {
	ID        int     `json:"id,omitempty"`
	NickName  string  `json:"nickname,omitempty"`
	Image     *string `json:"image,omitempty"`
	PhoneNum  string  `json:"phone_num,omitempty"`
	HuaweiID  *string `json:"huawei_id,omitempty"`
	Password  string  `json:"password,omitempty"`
	PushToken string  `json:"push_token,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}
