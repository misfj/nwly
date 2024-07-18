package service

import "github.com/gin-gonic/gin"

type ServiceRouter struct{}

func (serviceRouter *ServiceRouter) InitServiceRouter(router *gin.RouterGroup) {}
