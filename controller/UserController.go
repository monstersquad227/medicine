package controller

import (
	"github.com/gin-gonic/gin"
	"medicine/model"
	"medicine/service"
	"medicine/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService service.UserInterface
}

func (ctrl *UserController) UserLogin(c *gin.Context) {
	req := &model.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.UserService.UserLogin(req.PhoneNum, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
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

func (ctrl *UserController) UserUpdate(c *gin.Context) {
	req := &model.User{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	req.ID = id

	err = ctrl.UserService.UserUpdate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (ctrl *UserController) UserUpdateNickname(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	req := &model.User{}
	if err = c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.UserService.UpdateNickname(id, req.NickName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *UserController) UserUpdatePhone(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	req := &model.User{}
	if err = c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.UserService.UpdatePhone(id, req.PhoneNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
