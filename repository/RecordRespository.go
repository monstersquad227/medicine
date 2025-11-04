package repository

import "medicine/model"

type RecordRepository struct{}

func (repo *RecordRepository) List(userId int) ([]*model.RecordModel, error) {
	query := "SELECT id, medicine_name, actual_time, memo, status, is_checked " +
		"FROM medicine_plan_record WHERE user_id = ? "
	rows, err := MysqlClient.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.RecordModel, 0)
	for rows.Next() {
		obj := model.RecordModel{}
		err = rows.Scan(&obj.ID, &obj.MedicineName, &obj.ActualTime, &obj.Memo, &obj.Status, &obj.IsChecked)
		if err != nil {
			return nil, err
		}
		data = append(data, &obj)
	}
	return data, nil
}

func (repo *RecordRepository) Create(record *model.RecordModel) (int64, error) {
	query := "INSERT " +
		"INTO medicine_plan_record(user_id, plan_id, actual_time, medicine_name, memo, status) " +
		"VALUES (?, ?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, record.UserID, record.PlanID, record.ActualTime, record.MedicineName, record.Memo, record.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *RecordRepository) Update(record *model.RecordModel) (int64, error) {
	query := "UPDATE medicine_plan_record " +
		"SET actual_time = ?, is_checked = ?, status = ? " +
		"WHERE plan_id = ? AND user_id = ? AND id = ? "
	result, err := MysqlClient.Exec(query, record.ActualTime, record.IsChecked, record.Status, record.PlanID, record.UserID, record.ID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (repo *RecordRepository) GetActualTimeByPlanIDANDUserID(planID, userID int) (string, error) {
	query := "SELECT actual_time " +
		"FROM medicine_plan_record " +
		"WHERE plan_id = ? AND user_id = ? "
	var actualTime string
	err := MysqlClient.QueryRow(query, planID, userID).Scan(&actualTime)
	if err != nil {
		return "", err
	}
	return actualTime, nil
}
