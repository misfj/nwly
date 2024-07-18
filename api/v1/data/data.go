package data

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DataApi struct{}

func (dataApi *DataApi) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
