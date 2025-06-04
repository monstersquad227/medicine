package repository

import (
	"medicine/model"
)

type PlanRepository struct{}

func (repo *PlanRepository) ListPlan(id int) ([]*model.CourseAndPlan, error) {
	query := "SELECT " +
		"    mc.medicine_name, " +
		"    mc.medicine_timing, " +
		"    mp.amount, " +
		"    mp.type, " +
		"    mp.plan_time, " +
		"    mp.id as plan_id " +
		"FROM " +
		"	medicine_course mc " +
		"LEFT JOIN " +
		"   medicine_plan mp ON mp.medicine_id = mc.id " +
		"WHERE " +
		"    mc.user_id = ? " +
		"ORDER BY " +
		"    mc.id ASC, mp.plan_time ASC;"
	rows, err := MysqlClient.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.CourseAndPlan, 0)
	for rows.Next() {
		obj := model.CourseAndPlan{}
		if err = rows.Scan(&obj.MedicineName, &obj.MedicineTiming, &obj.Amount, &obj.Type, &obj.PlanTime, &obj.PlanID); err != nil {
			return nil, err
		}
		data = append(data, &obj)
	}
	return data, nil
}

func (repo *PlanRepository) CreatePlan(r *model.CourseAndPlan) (int64, error) {
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
