package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/labstack/gommon/log"
	"fmt"
)

var origin = "http://localhost/"
var url = "ws://localhost:9999/ws"

func GetClient(ctx *gin.Context) {
	ctx.HTML(200, "index", gin.H{"Title":"Paybox: Mock device program"})
}

func BillPost(ctx *gin.Context) {
	bill := ctx.DefaultPostForm("bill", "100")
	ctx.JSON(200, gin.H{
		"status": "posted",
		"message": bill,
	})
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte(bill)
	err = conn.WriteJSON(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send:%s\n", msg)

	var msgRead = make([]byte, 512)
	err = conn.ReadJSON(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msgRead)
	ctx.JSON(200, msgRead)
}