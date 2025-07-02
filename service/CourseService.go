package service

import (
	"medicine/model"
	"medicine/repository"
	"strings"
)

type CourseService struct {
	CourseRepo *repository.CourseRepository
	UserRepo   *repository.UserRepository
	PlanRepo   *repository.PlanRepository
	RecordRepo *repository.RecordRepository
}

func (svc *CourseService) List(phone string) ([]*model.CourseAndPlan, error) {
	userID, err := svc.UserRepo.GetUserIDByPhoneNum(phone)
	if err != nil {
		return nil, err
	}
	//return svc.CourseRepo.ListCourse(userID)
	return svc.CourseRepo.ListCourseV2(userID)
}

func (svc *CourseService) Create(course *model.CourseAndPlan) (int64, error) {
	dateTimes := strings.Split(course.CourseStartTime, " ")
	date := dateTimes[0]

	medicineID, err := svc.CourseRepo.CreateCourse(course)
	if err != nil {
		return 0, err
	}

	course.MedicineID = int(medicineID)
	planInsertCount := 0
	for i := 0; i < len(course.CourseStartTimes); i++ {
		course.PlanTime = course.CourseStartTimes[i]
		planID, err := svc.PlanRepo.CreatePlan(course)
		if err != nil {
			return 0, err
		}
		record := &model.RecordModel{
			UserID:       course.UserId,
			PlanID:       int(planID),
			MedicineName: course.MedicineName,
			ActualTime:   date + " " + course.CourseStartTimes[i],
			Memo:         nil,
			IsChecked:    0,
			Status:       0,
		}
		_, err = svc.RecordRepo.Create(record)
		if err != nil {
			return 0, err
		}
		planInsertCount++
	}
	//planID, err := svc.PlanRepo.CreatePlan(course)
	//if err != nil {
	//	return 0, err
	//}
	return int64(planInsertCount), nil
}

func (svc *CourseService) Update(course *model.Course) (int64, error) {
	return svc.CourseRepo.UpdateCourse(course)
}

func (svc *CourseService) Modify(course *model.CourseAndPlan) (int64, error) {
	return svc.CourseRepo.UpdateCourseV2(course)
}

func (svc *CourseService) Delete(course *model.Course) (int64, error) {
	return svc.CourseRepo.RemoveCourse(course)
}
