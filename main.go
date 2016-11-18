package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mrtomyum/paybox_device/ctrl"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("view/**/*.tpl")
	//r.Static("/public", "./view/public")
	app := ctrl.Router(r)
	app.Run(":9999")
}