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

	api.GET("plan/:id", ctrl.ListPlanV2)
	//api.PUT("plan/:id", ctrl.UpdatePlan)
	//api.PATCH("plan/:id", ctrl.PatchPlan)

}
