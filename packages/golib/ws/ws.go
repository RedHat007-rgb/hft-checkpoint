package ws

import (
	"log"

	"github.com/gorilla/websocket"
)




func WebSocketConnection(WsUrl string) *websocket.Conn{
	conn,_,err:=websocket.DefaultDialer.Dial(WsUrl,nil)
	if err!=nil {
		log.Fatalf("error while dialing ws.%v",err)
	}
	return conn
}