package main

import (
	"fmt"
	"github.com/GoldenLeeK/go-gin-blog/models"
	"github.com/GoldenLeeK/go-gin-blog/pkg/gredis"
	"github.com/GoldenLeeK/go-gin-blog/pkg/logging"
	"log"
	"syscall"

	"github.com/GoldenLeeK/go-gin-blog/pkg/setting"
	"github.com/GoldenLeeK/go-gin-blog/routers"
	"github.com/fvbock/endless"
)

func main() {

	setting.Setup()
	models.Setup()
	logging.Setup()
	err := gredis.Setup()
	if err != nil {
		panic(fmt.Sprintf("redis server error : %v", err))
	}

	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
