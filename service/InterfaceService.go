package service

import "medicine/model"

type UserInterface interface {
	UserLoginV22(code string) (map[string]interface{}, error)
	//UserLogin(phone, password string) (map[string]interface{}, error)
	//UserUpdate(user *model.User) error
	//UpdateNickname(id int, nickname string) (int64, error)
	//UpdatePhone(id int, phoneNum string) (int64, error)
}

type CourseInterface interface {
	List(phone string) ([]*model.CourseAndPlan, error)
	Create(medicineCourse *model.CourseAndPlan) (int64, error)
	Update(course *model.Course) (int64, error)
	Modify(course *model.CourseAndPlan) (int64, error)
	Delete(course *model.Course) (int64, error)
}

type PlanServiceInterface interface {
	List(userID int) ([]*model.CourseAndPlan, error)
	ListV2(userID int, date string) ([]*model.CourseAndPlan, error)
	Create(plan *model.Plan) (int64, error)
	//Update(plan *model.Plan) (int64, error)
}

type RecordServiceInterface interface {
	Fetch(userID int) ([]*model.RecordModel, error)
	FetchV2(userID int) (interface{}, error)
	Update(record *model.RecordModel) (int64, error)
}
