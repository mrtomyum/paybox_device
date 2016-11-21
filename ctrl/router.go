package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", GetClient)
	r.POST("/bill", BillPost)
	r.GET("/ws", func(c *gin.Context) {
		ws(c.Writer, c.Request)
	})
	return r
}

type msgStatus int
const (
	OK msgStatus = iota
	ERROR
	FAIL
)

type Msg struct {
	Status msgStatus `json:"status"`
	Data interface{} `json:"data"`
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	msg := Msg{}
	go func() {
		for {
			err := conn.ReadJSON(&msg)
			if err != nil {
				break
			}
			// Do something and assign new data to msg.Data
			conn.WriteJSON(&msg)
		}
	}()
}