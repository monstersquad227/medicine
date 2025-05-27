package service

import (
	"database/sql"
	"errors"
	"medicine/model"
	"medicine/repository"
	"medicine/utils"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (svc *UserService) UserLogin(phone, password string) (map[string]interface{}, error) {
	user, err := svc.UserRepo.GetUserInfo(phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 密码校验
	plaintext, err := utils.DecryptAESGCM(user.Password)
	if err != nil {
		return nil, err
	}
	if plaintext != password {
		return nil, errors.New("密码错误")
	}

	token, err := utils.GenerateToken(user.PhoneNum)
	user.PhoneNum = utils.HidePhoneNumber(user.PhoneNum)
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	return resp, nil
}

func (svc *UserService) UserUpdate(user *model.User) error {
	cipherPassword, err := utils.EncryptAESGCM(user.Password)
	if err != nil {
		return err
	}
	user.Password = cipherPassword

	ok, err := svc.UserRepo.UserUpdate(user)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("更新错误")
	}
	return nil
}
