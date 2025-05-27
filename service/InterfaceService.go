package service

import "medicine/model"

type UserInterface interface {
	UserLogin(phone, password string) (map[string]interface{}, error)
	UserUpdate(user *model.User) error
}

type CourseInterface interface {
	List(phone string) ([]*model.Course, error)
	Create(medicineCourse *model.CourseAndPlan) (int64, error)
	Update(course *model.Course) (int64, error)
	PatchCourseStatus(course *model.Course) (int64, error)
}

type PlanServiceInterface interface {
	Create(plan *model.Plan) (int64, error)
	Update(plan *model.Plan) (int64, error)
}
