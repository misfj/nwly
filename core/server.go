package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nwyl/global"
	"nwyl/initilalize"
	"strconv"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
}
func Start() {
	//todo 启动定时任务
	r := initilalize.Routers()
	s := initServer(global.Config.Server.Path+":"+strconv.Itoa(global.Config.Server.Port), r)
	global.Logger.Info(fmt.Sprintf("Server started:%s\n", global.Config.Server.Path+":"+strconv.Itoa(global.Config.Server.Port)))
	err := s.ListenAndServe()
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
}
