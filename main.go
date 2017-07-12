package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mrtomyum/paybox_device/ctrl"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("view/**/*")
	//r.Static("/public", "./view/public")
	//r.Static("/css", "./view/public/css")
	//r.Static("/js", "./view/public/js")
	app := ctrl.Router(r)
	app.Run(":9999")
}