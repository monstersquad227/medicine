package service

import (
	"medicine/model"
	"medicine/repository"
)

type PlanService struct {
	PlanRepo *repository.PlanRepository
}

func (svc *PlanService) Create(plan *model.Plan) (int64, error) {
	return svc.PlanRepo.CreatePlan(plan)
}

func (svc *PlanService) Update(plan *model.Plan) (int64, error) {
	return svc.PlanRepo.UpdatePlan(plan)
}
