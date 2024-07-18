package service

import "nwyl/service/app"
import "nwyl/service/system"

type ServiceGroup struct {
	SystemGroup system.ServiceGroup
	AppGroup    app.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
