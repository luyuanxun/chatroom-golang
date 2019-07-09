package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type Message struct {
	Type    string    `json:"type"`
	Content string `json:"content"`
}

//map记录所有ws客户端
var (
	clients      = make(map[string]*websocket.Conn)
	broadcastMsg = make(chan []byte, 100)
)

// gin 不支持websocket，需升级http请求为webSocket协议
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//webSocket handle
func webSocket(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	//生成一个clientId
	uuid1, _ := uuid.NewV4()
	clientId := uuid1.String()
	defer func() {
		//连接断开，删除无效client
		delete(clients, clientId)
		ws.Close()
	}()

	//记录每一个新连接
	clients[clientId] = ws

	for {
		//读取ws中的数据
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		//加入广播消息
		broadcastMsg <- msg
	}
}

//广播(群发)
func broadcast() {
	for {
		v, ok := <-broadcastMsg
		if !ok {
			break
		}

		go func() {
			for id, client := range clients {
				//tip：id可以判断不给谁发消息
				if err := client.WriteMessage(websocket.TextMessage, v); err != nil {
					//发送失败，客户端异常（断线...）
					delete(clients, id)
				}
			}
		}()
	}
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ws", webSocket)
	r.StaticFile("/", "./index.html")

	//启动一个协程处理广播消息
	go broadcast()
	r.Run("0.0.0.0:9090")
}
