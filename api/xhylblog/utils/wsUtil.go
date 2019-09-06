package utils

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)
var (
	MAX_CONN_CHANNEL_CACHE = 10000
)

func init() {
	if wsMaxCount, err :=beego.AppConfig.Int("MAX_CONN_CHANNEL_CACHE");err==nil{
		MAX_CONN_CHANNEL_CACHE=wsMaxCount
	}
}


//配置websocket最大创建数
var MaxConnChannel = make(chan struct{}, MAX_CONN_CHANNEL_CACHE)

// http升级websocket协议的配置
var WsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端读写消息
type wsMessage struct {
	messageType int
	data []byte
}

// 客户端连接
type WsConnection struct {
	WsSocket *websocket.Conn // 底层websocket
	InChan chan *wsMessage	// 读队列
	OutChan chan *wsMessage // 写队列

	Mutex sync.Mutex	// 避免重复关闭管道
	IsClosed bool
	CloseChan chan byte  // 关闭通知，写协程和读协程相互通知关闭连接
	MaxConnChannel chan struct{}  // 可创建的websocket最大连接数
}

func (wsConn *WsConnection)wsReadLoop() {
	for {
		// 读一个message
		msgType, data, err := wsConn.WsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		req := &wsMessage{
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

func (wsConn *WsConnection)wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <- wsConn.OutChan:
			// 写给websocket
			if err := wsConn.WsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
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

func (wsConn *WsConnection)procLoop() {
	//如果客户端创建websocket连接超过最大数量则阻塞
	MaxConnChannel <- struct{}{}

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
			fmt.Println(string(msg.data))
			err = wsConn.WsWrite(msg.messageType, msg.data)
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
func WsHandler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	WsSocket, err := WsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}
	wsConn := &WsConnection{
		WsSocket: WsSocket,
		InChan: make(chan *wsMessage, 1000),
		OutChan: make(chan *wsMessage, 1000),
		CloseChan: make(chan byte),
		IsClosed: false,
		MaxConnChannel: MaxConnChannel ,
	}

	// 处理器
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

func (wsConn *WsConnection)WsWrite(messageType int, data []byte) error {
	select {
	case wsConn.OutChan <- &wsMessage{messageType, data,}:
	case <- wsConn.CloseChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *WsConnection)WsRead() (*wsMessage, error) {
	select {
	case msg := <- wsConn.InChan:
		return msg, nil
	case <- wsConn.CloseChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *WsConnection)WsClose() {
	wsConn.WsSocket.Close()

	wsConn.Mutex.Lock()
	defer wsConn.Mutex.Unlock()

	if !wsConn.IsClosed {
		wsConn.IsClosed = true
		close(wsConn.CloseChan)
	}

	<- MaxConnChannel
}

//func main() {
//	http.HandleFunc("/ws", wsHandler)
//	http.ListenAndServe("0.0.0.0:7777", nil)
//}