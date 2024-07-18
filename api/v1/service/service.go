package service

import "github.com/gin-gonic/gin"
import "net/http"

type ServiceApi struct{}

func (serviceApi *ServiceApi) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
