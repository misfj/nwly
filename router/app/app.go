package app

import (
	"github.com/gin-gonic/gin"
	v1 "nwyl/api/v1"
)

type AppRouter struct{}

func (e *AppRouter) InitAppRouter(Router *gin.RouterGroup) {
	//customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouter := Router.Group("customer")

	//customerRouterWithoutRecord := Router.Group("customer")
	exaCustomerApi := v1.ApiGroupApp.AppGroup.AppApi
	{
		customerRouter.POST("customer", exaCustomerApi.Create) // 创建客户
		//customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		//customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		//customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		//customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
	//return customerRouter
}
