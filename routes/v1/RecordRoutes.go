package v1

import (
	"github.com/gin-gonic/gin"
	"medicine/controller"
	"medicine/repository"
	"medicine/service"
)

func RecordRegister(api *gin.RouterGroup) {

	ctrl := &controller.RecordController{
		RecordService: &service.RecordService{
			RecordRepository: &repository.RecordRepository{},
		},
	}
	api.GET("record/:id", ctrl.ListRecord)
	api.PUT("/record/:id", ctrl.UpdateRecord)
}
