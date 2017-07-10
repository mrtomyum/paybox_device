package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
	"time"
	"github.com/mrtomyum/paybox_dev/model"
)

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
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Println(err)
	}
	defer c.Close()
	msg := model.Message{}

	go func() {
		defer c.Close()
		for {
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("receive: %s", msg)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	//msg = Msg{Status:OK, Message:bill}
	//byteMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error on json.Marshal:", err)
	}
	for t := range ticker.C {
		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
		if err != nil {
			log.Println("write:", err)
			break
		}
		fmt.Printf("Send:%s\n", msg)
		ctx.HTML(200,"reply", msg)
	}


}