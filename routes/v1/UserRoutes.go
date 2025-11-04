package v1

import (
	"github.com/gin-gonic/gin"
	"medicine/controller"
	"medicine/repository"
	"medicine/service"
)

func UserRegister(api *gin.RouterGroup) {

	ctrl := &controller.UserController{
		UserService: &service.UserService{
			UserRepo: &repository.UserRepository{},
		},
	}

	api.POST("user/loginV2", ctrl.UserLoginV2)
	api.PUT("user", ctrl.UserPushToken)
}
