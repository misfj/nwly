package v1

import (
	"nwyl/api/v1/app"
	"nwyl/api/v1/data"
	"nwyl/api/v1/node"
	"nwyl/api/v1/service"
	"nwyl/api/v1/system"
	"nwyl/api/v1/user"
)

type ApiGroup struct {
	AppGroup     app.AppGroup
	DataGroup    data.DataGroup
	NodeGroup    node.NodeGroup
	ServiceGroup service.ServiceGroup
	SystemGroup  system.ServiceGroup
	UserGroup    user.UserGroup
}

var ApiGroupApp = new(ApiGroup)
