package app

import "nwyl/service"

// AppGroup 应用管理
type AppGroup struct {
	AppApi
	FileApi
}

var (
	appService = service.ServiceGroupApp.AppGroup.AppService
)
