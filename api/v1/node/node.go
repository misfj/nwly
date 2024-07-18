package node

import "github.com/gin-gonic/gin"
import "net/http"

type NodeApi struct{}

func (nodeApi *NodeApi) Create(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
