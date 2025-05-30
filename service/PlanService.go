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

func (svc *PlanService) Create(plan *model.Plan) (int64, error) {
	//return svc.PlanRepo.CreatePlan(plan)
	return 0, nil
}

func (svc *PlanService) Update(plan *model.Plan) (int64, error) {
	return svc.PlanRepo.UpdatePlan(plan)
}
