package repository

import "medicine/model"

type PlanRepository struct{}

func (repo *PlanRepository) CreatePlan(r *model.CourseReq) (int64, error) {
	query := "INSERT " +
		"INTO medicine_plan(medicine_id, amount, type, plan_time) " +
		"VALUES (?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, r.MedicineID, r.Amount, r.Type, r.PlanTime)
	if err != nil {
		return 0, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (repo *PlanRepository) UpdatePlan(plan *model.Plan) (int64, error) {
	return 0, nil
}
