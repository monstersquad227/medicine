package service

import (
	"fmt"
	"medicine/model"
	"medicine/repository"
)

type CourseService struct {
	CourseRepo *repository.CourseRepository
	UserRepo   *repository.UserRepository
	PlanRepo   *repository.PlanRepository
}

func (svc *CourseService) List(phone string) ([]*model.Course, error) {
	userID, err := svc.UserRepo.GetUserIDByPhoneNum(phone)
	if err != nil {
		return nil, err
	}
	return svc.CourseRepo.ListCourse(userID)
}

func (svc *CourseService) Create(course *model.CourseAndPlan) (int64, error) {
	medicineID, err := svc.CourseRepo.CreateCourse(course)
	fmt.Println(medicineID)
	if err != nil {
		return 0, err
	}
	course.MedicineID = int(medicineID)
	planID, err := svc.PlanRepo.CreatePlan(course)
	if err != nil {
		return 0, err
	}
	return planID, nil
}

func (svc *CourseService) Update(course *model.Course) (int64, error) {
	return svc.CourseRepo.UpdateCourse(course)
}

func (svc *CourseService) PatchCourseStatus(course *model.Course) (int64, error) {
	return svc.CourseRepo.UpdateCourseStatusByID(course)
}
