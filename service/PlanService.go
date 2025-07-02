package service

import (
	"medicine/model"
	"medicine/repository"
)

type PlanService struct {
	PlanRepo *repository.PlanRepository
}

func (svc *PlanService) List(userID int) ([]*model.CourseAndPlan, error) {
	return svc.PlanRepo.ListPlan(userID)
}

func (svc *PlanService) ListV2(userID int, date string) ([]*model.CourseAndPlan, error) {
	startTime := date + " 00:00:00"
	endTime := date + " 23:59:59"
	return svc.PlanRepo.ListPlanV2(userID, startTime, endTime)
}

func (svc *PlanService) Create(plan *model.Plan) (int64, error) {
	//return svc.PlanRepo.CreatePlan(plan)
	return 0, nil
}

//func (svc *PlanService) Update(plan *model.Plan) (int64, error) {
//	return svc.PlanRepo.UpdatePlan(plan)
//}
