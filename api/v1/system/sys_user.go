package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemApi struct{}

// CreateApi
// @Tags      SysApi
// @Summary   创建基础api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建基础api"
// @Router    /api/createApi [post]
func (s *SystemApi) CreateApi(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

// CreateApi2
// @Tags      SysApi1
// @Summary   创建基础api1
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      system.SysApi                  true  "api路径, api中文描述, api组, 方法"
// @Success   200   {object}  response.Response{msg=string}  "创建基础api"
// @Router    /api/createApi2 [post]
func (s *SystemApi) CreateApi2(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
