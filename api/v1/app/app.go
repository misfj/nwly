package app

import (
	"github.com/gin-gonic/gin"
)
import "net/http"

type AppApi struct{}

func (e *AppApi) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")

}
func (e *AppApi) CreateFunc() {
	////appService.AppInBlacklist()
	//fmt.Println("create")
}
