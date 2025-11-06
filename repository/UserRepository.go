package repository

import (
	"medicine/model"
)

type UserRepository struct{}

func (repo *UserRepository) CreateUser(user *model.User) (int64, error) {
	query := "INSERT " +
		"INTO user(phone_num, huawei_id, password) " +
		"VALUES (?, ?, ?) " +
		"ON DUPLICATE KEY UPDATE " +
		"password = VALUES(password) "
	exec, err := MysqlClient.Exec(query, user.PhoneNum, user.HuaweiID, user.Password)
	if err != nil {
		return 0, err
	}
	return exec.LastInsertId()
}

func (repo *UserRepository) GetUserById(id int64) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, nickname, image, phone_num, huawei_id, password, created_at, updated_at " +
		"FROM user WHERE id = ?"
	err := MysqlClient.QueryRow(query, id).Scan(
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

func (repo *UserRepository) GetUserInfo(phone string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, nickname, image, phone_num, huawei_id, password, created_at, updated_at " +
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

func (repo *UserRepository) UpdateNickname(id int, nickname string) (int64, error) {
	query := "UPDATE user " +
		"SET nickname = ? WHERE id = ?"
	exec, err := MysqlClient.Exec(query, nickname, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (repo *UserRepository) UpdatePhone(id int, phoneNum string) (int64, error) {
	query := "UPDATE user " +
		"SET phone_num = ? WHERE id = ?"
	exec, err := MysqlClient.Exec(query, phoneNum, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (repo *UserRepository) UpdatePushToken(phone, pushToken string, isEnabled int) (int64, error) {
	query := "UPDATE user " +
		"SET push_token = ?, notify_enabled = ? " +
		"WHERE phone_num = ?"
	exec, err := MysqlClient.Exec(query, pushToken, isEnabled, phone)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}
