package router

import (
	"nwyl/router/app"
	"nwyl/router/data"
	"nwyl/router/node"
	"nwyl/router/service"
	"nwyl/router/system"
	"nwyl/router/user"
)

type RouterGroup struct {
	AppRouter     app.AppRouter
	DataRouter    data.DataRouter
	NodeRouter    node.NodeRouter
	ServiceRouter service.ServiceRouter
	SystemRouter  system.SystemRouter
	UserRouter    user.UserRouter
}

var RouterGroupApp = new(RouterGroup)
