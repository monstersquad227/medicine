package service

import (
	"fmt"
	"medicine/model"
	"medicine/repository"
	"medicine/utils"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

//func (svc *UserService) UserLogin(phone, password string) (map[string]interface{}, error) {
//	user, err := svc.UserRepo.GetUserInfo(phone)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, errors.New("用户不存在")
//		}
//		return nil, err
//	}
//
//	// 密码校验
//	plaintext, err := utils.DecryptAESGCM(user.Password)
//	if err != nil {
//		return nil, err
//	}
//	if plaintext != password {
//		return nil, errors.New("密码错误")
//	}
//
//	token, err := utils.GenerateToken(user.PhoneNum)
//	//user.PhoneNum = utils.HidePhoneNumber(user.PhoneNum)
//	if err != nil {
//		return nil, err
//	}
//	resp := map[string]interface{}{
//		"token": token,
//		"user":  user,
//	}
//	return resp, nil
//}

func (svc *UserService) UserLoginV22(code string) (map[string]interface{}, error) {

	accessToken, err := utils.GetHuaweiAccessToken(code)
	if err != nil || accessToken == "" {
		return nil, err
	}

	UnionID, LoginMobileNumber, err := utils.GetHuaweiUserInfo(accessToken)
	if err != nil || UnionID == "" || LoginMobileNumber == "" {
		return nil, err
	}

	encryptToken, err := utils.EncryptAESGCM(accessToken)
	if err != nil {
		return nil, err
	}

	user := model.User{}
	user.HuaweiID = &UnionID
	user.Password = encryptToken
	user.PhoneNum = LoginMobileNumber

	insertID, err := svc.UserRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	userInfo, err := svc.UserRepo.GetUserById(insertID)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.PhoneNum)
	fmt.Println("token: ", token)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"token": token,
		"user":  userInfo,
	}

	return resp, nil
}

func (svc *UserService) UserUpdatePushToken(phone, pushToken string) (int64, error) {
	return svc.UserRepo.UpdatePushToken(phone, pushToken)
}

//func (svc *UserService) UserUpdate(user *model.User) error {
//	cipherPassword, err := utils.EncryptAESGCM(user.Password)
//	if err != nil {
//		return err
//	}
//	user.Password = cipherPassword
//
//	ok, err := svc.UserRepo.UserUpdate(user)
//	if err != nil {
//		return err
//	}
//	if !ok {
//		return errors.New("更新错误")
//	}
//	return nil
//}
//
//func (svc *UserService) UpdateNickname(id int, nickname string) (int64, error) {
//	return svc.UserRepo.UpdateNickname(id, nickname)
//}
//
//func (svc *UserService) UpdatePhone(id int, phoneNum string) (int64, error) {
//	return svc.UserRepo.UpdatePhone(id, phoneNum)
//}
