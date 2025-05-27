package v1

import (
	"github.com/gin-gonic/gin"
	"medicine/controller"
	"medicine/repository"
	"medicine/service"
)

func PlanRegister(api *gin.RouterGroup) {
	ctrl := &controller.PlanController{
		PlanService: &service.PlanService{
			PlanRepo: &repository.PlanRepository{},
		},
	}

	//api.POST("plan", ctrl.CreatePlan) 创建用药计划不对外
	api.PUT("plan/:id", ctrl.UpdatePlan)
	api.PATCH("plan/:id", ctrl.PatchPlan)

}
