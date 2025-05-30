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
	data := make([]*model.RecordModel, 0)
	if rows.Next() {
		var obj model.RecordModel
		err = rows.Scan(&obj.ID, &obj.MedicineName, &obj.ActualTime, &obj.Memo, &obj.Status)
		if err != nil {
			return nil, err
		}
		data = append(data, &obj)
	}
	return data, nil
}
