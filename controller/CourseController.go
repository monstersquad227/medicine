package controller

import (
	"github.com/gin-gonic/gin"
	"medicine/model"
	"medicine/service"
	"medicine/utils"
	"net/http"
	"strconv"
)

type CourseController struct {
	CourseService service.CourseInterface
}

func (ctrl *CourseController) ListCourse(c *gin.Context) {
	phone, _ := c.Get("account")
	result, err := ctrl.CourseService.List(phone.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	req := &model.CourseAndPlan{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}

	result, err := ctrl.CourseService.Create(req)
	if err != nil {
		c.JSON(http.StatusCreated, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusCreated, utils.Success(result))
}

func (ctrl *CourseController) UpdateCourse(c *gin.Context) {
	req := &model.Course{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	req.ID = id
	result, err := ctrl.CourseService.Update(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *CourseController) UpdateCourseV2(c *gin.Context) {
	req := &model.CourseAndPlan{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}

	result, err := ctrl.CourseService.Modify(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *CourseController) DeleteCourse(c *gin.Context) {
	req := &model.Course{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}

	result, err := ctrl.CourseService.Delete(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}

func (ctrl *CourseController) RestoreCourse(c *gin.Context) {
	req := &model.Course{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(1, err.Error(), err))
		return
	}
	result, err := ctrl.CourseService.Restore(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(1, err.Error(), err))
		return
	}
	c.JSON(http.StatusOK, utils.Success(result))
}
