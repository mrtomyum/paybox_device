package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Router(r *gin.Engine) *gin.Engine {
	r.LoadHTMLGlob("view/**/*")
	//r.Static("/public", "./view/public")
	//r.Static("/css", "./view/public/css")
	//r.Static("/js", "./view/public/js")

	//r.GET("/", GetClient)
	//r.POST("/bill", BillPost)
	//r.POST("/coin", CoinPost)
	r.GET("/", func(c *gin.Context) {
		ServDev(c.Writer, c.Request)
	})
	return r
}

