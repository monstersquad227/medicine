package v1

import (
	"github.com/gin-gonic/gin"
	"medicine/controller"
	"medicine/repository"
	"medicine/service"
)

func CourseRegister(api *gin.RouterGroup) {
	ctrl := &controller.CourseController{
		CourseService: &service.CourseService{
			CourseRepo: &repository.CourseRepository{},
		},
	}

	api.GET("course", ctrl.ListCourse)
	api.POST("/course", ctrl.CreateCourse)
	api.PUT("/course/:id", ctrl.UpdateCourse)
	api.PATCH("/course/:id", ctrl.PatchCourseStatus)
}
