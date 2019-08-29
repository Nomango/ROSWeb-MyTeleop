package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/nomango/ROSWeb-MyTeleop/modules/settings"
	"github.com/nomango/ROSWeb-MyTeleop/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	settings.Init()
	routers.Setup(engine)

	server := http.Server{Handler: engine, Addr: ":" + settings.Port}
	if err := server.ListenAndServe(); err != nil {
		logrus.Panicln("Start Gin server failed", err)
	}
}
