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
		"    mp.id as plan_id, " +
		"    IF(mpr.is_checked = 1, 1, 0) AS is_checked " +
		"FROM " +
		"	medicine_course mc " +
		"LEFT JOIN " +
		"   medicine_plan mp ON mp.medicine_id = mc.id " +
		"LEFT JOIN ( " +
		"	SELECT " +
		"		plan_id, " +
		"		MAX(is_checked) AS is_checked " +
		"	FROM " +
		"		medicine_plan_record " +
		"	GROUP BY " +
		"		plan_id" +
		") mpr ON mpr.plan_id = mp.id " +
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
		if err = rows.Scan(&obj.MedicineName, &obj.MedicineTiming, &obj.Amount, &obj.Type, &obj.PlanTime, &obj.PlanID, &obj.IsChecked); err != nil {
			return nil, err
		}
		data = append(data, &obj)
	}
	return data, nil
}

func (repo *PlanRepository) ListPlanV2(id, status int, startTime, endTime string) ([]*model.CourseAndPlan, error) {
	query := "SELECT " +
		"    mp.id AS plan_id, " +
		"    mc.medicine_name, " +
		"    mc.medicine_type, " +
		"    mc.medicine_timing, " +
		"    mp.amount, " +
		"    mp.type, " +
		"    mp.plan_time, " +
		"    IFNULL(mpr.is_checked, 0) AS is_checked, " +
		"    mpr.id AS record_id " +
		"FROM " +
		"    medicine_plan mp " +
		"JOIN " +
		"    medicine_course mc ON mp.medicine_id = mc.id " +
		"LEFT JOIN " +
		"    medicine_plan_record mpr " +
		"    ON mpr.plan_id = mp.id " +
		"    AND mpr.user_id = ? " +
		"WHERE " +
		"    mc.user_id = ? " +
		"    AND (" +
		"        (? = 0 AND mc.status = 0) OR " +
		"        (? = 1 AND mc.status IN (0, 1)) " +
		"    ) " +
		"    AND ( " +
		"        mpr.id IS NULL " +
		"        OR mpr.actual_time BETWEEN ? AND ? " +
		"    ) " +
		"ORDER BY " +
		"    mp.plan_time ASC, mp.id ASC;"
	rows, err := MysqlClient.Query(query, id, id, status, status, startTime, endTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.CourseAndPlan, 0)
	for rows.Next() {
		obj := model.CourseAndPlan{}
		if err = rows.Scan(
			&obj.PlanID,
			&obj.MedicineName,
			&obj.MedicineType,
			&obj.MedicineTiming,
			&obj.Amount,
			&obj.Type,
			&obj.PlanTime,
			&obj.IsChecked,
			&obj.RecordID); err != nil {
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

func (repo *PlanRepository) GetPlanTimeByIdAndUserID(id int) (string, error) {
	query := "SELECT plan_time " +
		"FROM medicine_plan " +
		"WHERE id = ? "
	var planTime string
	err := MysqlClient.QueryRow(query, id).Scan(&planTime)
	if err != nil {
		return "", err
	}
	return planTime, nil
}

func (repo *PlanRepository) GetPlanIDsByCourseID(id int) ([]int64, error) {
	query := "SELECT id " +
		"FROM medicine_plan " +
		"WHERE medicine_id = ? "
	rows, err := MysqlClient.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var planIDs []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		planIDs = append(planIDs, id)
	}
	return planIDs, nil
}
