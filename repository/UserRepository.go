package repository

import (
	"medicine/model"
)

type UserRepository struct{}

func (repo *UserRepository) GetUserInfo(phone string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT * " +
		"FROM user WHERE phone_num = ?"
	err := MysqlClient.QueryRow(query, phone).Scan(
		&user.ID,
		&user.NickName,
		&user.Image,
		&user.PhoneNum,
		&user.HuaweiID,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetUserIDByPhoneNum(phone string) (int, error) {
	var id int
	query := "SELECT id " +
		"FROM user " +
		"WHERE phone_num = ?"
	err := MysqlClient.QueryRow(query, phone).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *UserRepository) UserUpdate(user *model.User) (bool, error) {
	query := "UPDATE user " +
		"SET nickname = ?, image = ?, phone_num = ?, huawei_id = ?, password = ? " +
		"WHERE id = ?"

	result, err := MysqlClient.Exec(query, user.NickName, user.Image, user.PhoneNum, user.HuaweiID, user.Password, user.ID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}
