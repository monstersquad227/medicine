package repository

import "medicine/model"

type PlanRepository struct{}

func (repo *PlanRepository) CreatePlan(plan *model.Plan) (int64, error) {
	return 0, nil
}

func (repo *PlanRepository) UpdatePlan(plan *model.Plan) (int64, error) {
	return 0, nil
}
