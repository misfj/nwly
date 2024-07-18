package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(router *gin.RouterGroup) {}
