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

	api.POST("user/login", ctrl.UserLogin)
	api.PUT("user/:id", ctrl.UserUpdate)
}
