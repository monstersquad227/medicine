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

	api.GET("course", ctrl.ListCourse)               // 有问题 需要改
	api.POST("/course", ctrl.CreateCourse)           // √
	api.PUT("/course/:id", ctrl.UpdateCourse)        // 有问题 需要改
	api.PATCH("/course/:id", ctrl.PatchCourseStatus) // √,只改 course 的 status 没问题
}
