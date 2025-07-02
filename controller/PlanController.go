package controller

import (
	"github.com/gin-gonic/gin"
	"medicine/model"
	"medicine/service"
	"medicine/utils"
	"net/http"
	"strconv"
)

type PlanController struct {
	PlanService service.PlanServiceInterface
}

func (ctrl *PlanController) ListPlan(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.PlanService.List(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *PlanController) ListPlanV2(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.PlanService.ListV2(id, c.Query("date"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *PlanController) CreatePlan(c *gin.Context) {
	req := &model.Plan{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.PlanService.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

//func (ctrl *PlanController) UpdatePlan(c *gin.Context) {
//	req := &model.Plan{}
//	if err := c.ShouldBindJSON(req); err != nil {
//		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
//		return
//	}
//	result, err := ctrl.PlanService.Update(req)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
//		return
//	}
//	c.JSON(http.StatusOK, utils.Success(result))
//}

//func (ctrl *PlanController) PatchPlan(c *gin.Context) {
//	req := &model.Plan{}
//	if err := c.ShouldBindJSON(req); err != nil {
//		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
//		return
//	}
//	c.JSON(http.StatusOK, utils.Success(nil))
//}
