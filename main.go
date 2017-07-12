package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtomyum/paybox_device/ctrl"
)

func main() {
	r := gin.Default()
	app := ctrl.Router(r)
	app.Run(":9999")
}