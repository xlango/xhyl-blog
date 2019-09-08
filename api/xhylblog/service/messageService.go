package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"xhylblog/utils"
)


type MessageServcie struct {

}



func (this *MessageServcie)wsReadLoop(wsConn utils.WsConnection) {
	for {
		// 读一个message
		msgType, data, err := wsConn.WsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := & utils.WsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.InChan <- req:
		case <- wsConn.CloseChan:
			goto closed
		}
	}
error:
	wsConn.WsClose()
closed:
}

func (this *MessageServcie)wsWriteLoop(wsConn utils.WsConnection) {
	for {
		select {
		// 取一个应答
		case msg := <- wsConn.OutChan:
			// 写给websocket
			if err := wsConn.WsSocket.WriteMessage(msg.MessageType, msg.Data); err != nil {
				goto error
			}
		case <- wsConn.CloseChan:
			goto closed
		}
	}
error:
	wsConn.WsClose()
closed:
}

func (this *MessageServcie)procLoop(wsConn utils.WsConnection) {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(10 * time.Second)
			if err := wsConn.WsWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				logs.Error("ws heartbeat fail")
				wsConn.WsClose()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	go func() {
		for {
			msg, err := wsConn.WsRead()
			if err != nil {
				logs.Error("ws read fail")
				break
			}
			fmt.Println(string(msg.Data))
			err = wsConn.WsWrite(msg.MessageType, msg.Data)
			if err != nil {
				logs.Error("ws write fail")
				break
			}
		}
	}()


}


/**
websocket消息处理
*/
func (this *MessageServcie)WsHandler(resp http.ResponseWriter, req *http.Request) {

	wsConn := utils.GetWsConn(resp,req)
	// 处理器
	go this.procLoop(wsConn)
	// 读协程
	go this.wsReadLoop(wsConn)
	// 写协程
	go this.wsWriteLoop(wsConn)
}