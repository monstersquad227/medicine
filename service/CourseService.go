package service

import (
	"medicine/model"
	"medicine/repository"
)

type CourseService struct {
	CourseRepo *repository.CourseRepository
	UserRepo   *repository.UserRepository
}

func (svc *CourseService) List(phone string) ([]*model.Course, error) {
	userID, err := svc.UserRepo.GetUserIDByPhoneNum(phone)
	if err != nil {
		return nil, err
	}
	return svc.CourseRepo.ListCourse(userID)
}

func (svc *CourseService) Create(course *model.Course) (int64, error) {
	return svc.CourseRepo.CreateCourse(course)
}

func (svc *CourseService) Update(course *model.Course) (int64, error) {
	return svc.CourseRepo.UpdateCourse(course)
}

func (svc *CourseService) PatchCourseStatus(course *model.Course) (int64, error) {
	return svc.CourseRepo.UpdateCourseStatusByID(course)
}
