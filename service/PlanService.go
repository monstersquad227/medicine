package service

import (
	"medicine/model"
	"medicine/repository"
	"time"
)

type PlanService struct {
	PlanRepo *repository.PlanRepository
}

func (svc *PlanService) List(userID int) ([]*model.CourseAndPlan, error) {
	return svc.PlanRepo.ListPlan(userID)
}

func (svc *PlanService) ListV2(userID int, date string) ([]*model.CourseAndPlan, error) {
	status := 0
	// 获取今天的日期（只保留年月日）
	today := time.Now().Format("2006-01-02")
	if date != today {
		status = 1
	}

	startTime := date + " 00:00:00"
	endTime := date + " 23:59:59"
	return svc.PlanRepo.ListPlanV2(userID, status, startTime, endTime)
}

func (svc *PlanService) Create(plan *model.Plan) (int64, error) {
	//return svc.PlanRepo.CreatePlan(plan)
	return 0, nil
}

//func (svc *PlanService) Update(plan *model.Plan) (int64, error) {
//	return svc.PlanRepo.UpdatePlan(plan)
//}
