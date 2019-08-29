package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/nomango/ros_web/controllers"
	"github.com/nomango/ros_web/modules/templates"
)

func Setup(engine *gin.Engine) {
	engine.Delims("{{", "}}")
	engine.SetFuncMap(templates.FuncMap)
	engine.LoadHTMLFiles(
		"./views/index.html",
	)
	engine.Static("/static", "./static")

	engine.GET("/", controllers.Index)
}
