package repository

import "medicine/model"

type RecordRepository struct{}

func (repo *RecordRepository) List(userId int) ([]*model.RecordModel, error) {
	query := "SELECT id, medicine_name, actual_time, memo, status " +
		"FROM medicine_plan_record WHERE user_id = ?"
	rows, err := MysqlClient.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]*model.RecordModel, 0)
	for rows.Next() {
		obj := model.RecordModel{}
		err = rows.Scan(&obj.ID, &obj.MedicineName, &obj.ActualTime, &obj.Memo, &obj.Status)
		if err != nil {
			return nil, err
		}
		data = append(data, &obj)
	}
	return data, nil
}

func (repo *RecordRepository) Create(record *model.RecordModel) (int64, error) {
	query := "INSERT " +
		"INTO medicine_plan_record(user_id, plan_id, medicine_name, memo, status) " +
		"VALUES (?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(query, record.UserID, record.PlanID, record.MedicineName, record.Memo, record.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
