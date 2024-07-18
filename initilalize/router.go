package initilalize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "nwyl/docs"
	"nwyl/middleware"
	"nwyl/router"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	Router.Use(middleware.Cors())
	privateRouter := Router.Group("")
	//publicRouter := Router.Group("")
	//应用管理的路由
	{
		router.RouterGroupApp.AppRouter.InitAppRouter(privateRouter)
		//appRouter := router.RouterGroupApp.AppRouter
	}

	//数据管理的路由
	{
		router.RouterGroupApp.DataRouter.InitDataRouter(privateRouter)
	}

	//节点管理的路由
	{
		router.RouterGroupApp.NodeRouter.InitNodeRouter(privateRouter)
	}

	//服务管理的路由
	{
		router.RouterGroupApp.ServiceRouter.InitServiceRouter(privateRouter)
	}

	//系统管理的路由
	{
		router.RouterGroupApp.SystemRouter.InitSystemRouter(privateRouter)
	}

	//用户管理的路由
	{
		router.RouterGroupApp.UserRouter.InitUserRouter(privateRouter)
	}

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return Router
}
