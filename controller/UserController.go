package controller

import (
	"github.com/gin-gonic/gin"
	"medicine/service"
	"medicine/utils"
	"net/http"
)

type UserController struct {
	UserService service.UserInterface
}

func (ctrl *UserController) UserLoginV2(c *gin.Context) {
	var req struct {
		AuthorizationCode string `json:"authorization_code"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.UserService.UserLoginV22(req.AuthorizationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *UserController) UserPushToken(c *gin.Context) {
	var req struct {
		PushToken string `json:"push_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}

	phone, exist := c.Get("account")
	if !exist {
		c.JSON(http.StatusBadRequest, utils.Error(1, "c.Get() Account not exist ", nil))
		return
	}

	result, err := ctrl.UserService.UserUpdatePushToken(phone.(string), req.PushToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, utils.Success(result))
}
