package service

import (
	"errors"
	"fmt"
	"medicine/model"
	"medicine/repository"
	"strings"
	"time"
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
	// Step 1: 更新课程状态
	if _, err := svc.CourseRepo.UpdateCourseStatusByID(course); err != nil {
		return 0, fmt.Errorf("更新课程状态失败: %w", err)
	}

	// Step 2: 获取课程关联的所有计划 ID
	ids, err := svc.PlanRepo.GetPlanIDsByCourseID(course.ID)
	if err != nil {
		return 0, fmt.Errorf("获取计划ID失败: %w", err)
	}
	if len(ids) == 0 {
		// 没有计划可删
		return 0, nil
	}

	// Step 3: 删除所有计划的今日记录
	date := time.Now().Format("2006-01-02")
	startTime := date + " 00:00:00"
	endTime := date + " 23:59:59"
	var deletedCount int64
	for _, id := range ids {
		ok, err := svc.RecordRepo.DeleteTodayRecordsByPlanID(id, startTime, endTime)
		if err != nil {
			return 0, fmt.Errorf("删除 plan_id=%d 的今日记录失败: %w", id, err)
		}
		if ok {
			deletedCount++
		}
	}

	return deletedCount, nil
}

func (svc *CourseService) Restore(course *model.Course) (int64, error) {
	// 恢复用药方案，状态置为 0
	_, err := svc.CourseRepo.UpdateCourseStatusByID(course)
	if err != nil {
		return 1, err
	}

	// 查看当天是否有打卡记录
	ids, err := svc.PlanRepo.GetPlanIDsByCourseID(course.ID)
	if err != nil {
		return 1, err
	}

	if len(ids) == 0 {
		return 1, errors.New("没有原方案计划")
	}

	today := time.Now().Format("2006-01-02")
	startTime := today + " 00:00:00"
	endTime := today + " 23:59:59"

	for _, id := range ids {
		ok, err := svc.RecordRepo.HasTodayRecordByPlanID(id, startTime, endTime)
		if err != nil {
			return 0, err
		}
		if ok == false {
			// 创建今天的用药记录
		}
	}

	return 0, nil
}
