package controller

import (
	"github.com/gin-gonic/gin"
	"medicine/model"
	"medicine/service"
	"medicine/utils"
	"net/http"
	"strconv"
)

type RecordController struct {
	RecordService service.RecordServiceInterface
}

func (ctrl *RecordController) List(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.RecordService.Fetch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *RecordController) Create(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	req := &model.RecordModel{}
	if err = c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.RecordService.Create(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
