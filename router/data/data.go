package data

import (
	"github.com/gin-gonic/gin"
	v1 "nwyl/api/v1"
)

type DataRouter struct{}

func (dataRouter *DataRouter) InitDataRouter(Router *gin.RouterGroup) {
	dtaRouter := Router.Group("data")
	{
		dtaRouter.POST("/kill", v1.ApiGroupApp.DataGroup.Create)
	}
}
