package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RedHat007-rgb/hft-checkpoint/packages/golib/redis"
	"github.com/RedHat007-rgb/hft-checkpoint/packages/golib/ws"
	"github.com/gorilla/websocket"
)

//websocket connection per each symbol
//publish data


type WsConnections struct{
	name string
	conn *websocket.Conn
}
var binanceWsConns=[]*WsConnections{}

func BinanceWsConnections(tokens []string)[]*WsConnections{
	for index,_ :=range tokens{
		url:="wss://stream.binance.com:9443/ws/"+tokens[index]+"@ticker"
		fmt.Println(url)
		WsName:="Ws"+tokens[index];
		 connection:=ws.WebSocketConnection(url)
		 c:=&WsConnections{
			name: WsName,
			conn: connection,
		 }
		 binanceWsConns=append(binanceWsConns,c)
		}
		return binanceWsConns
}


func main() {
	ctx := context.Background()
	tokens := []string{"btcusdt", "solusdt", "ethusdt"}
	redisClient := redis.NewConnection()
	BinanceWsConnections(tokens)

	for index := range binanceWsConns {
		wsConn := binanceWsConns[index].conn
		token := tokens[index]
		// start a goroutine per websocket
		go func(c *websocket.Conn, t string) {
			defer c.Close()
			channel := "binance." + t + ".ticker"
			for {
				_, msg, err := c.ReadMessage()
				if err != nil {
					log.Printf("error reading message for %s: %v", t, err)
					return
				}
				if err := redisClient.PublishMessages(ctx, channel, msg); err != nil {
					log.Printf("error while publishing: %v", err)
				}
			}
		}(wsConn, token)
	}

	// prevent main from exiting
	select {}
}



// func main(){
// 	ctx:=context.Background()
// 	tokens:=[]string{"btcusdt","solusdt","ethusdt"}
// 	redisClient:=redis.NewConnection()
// 	BinanceWsConnections(tokens)
// 	for index:=range binanceWsConns{
// 		wsConn:=binanceWsConns[index].conn
// 		defer wsConn.Close()
// 		_,msg,err:=wsConn.ReadMessage()
// 		if err!=nil{
// 			log.Fatalf("error reading message with error %v",err)
// 		}
// 		channel:="binance."+tokens[index]+".ticker"
// 		fmt.Println(channel)
// 		if err:=redisClient.PublishMessages(ctx,channel,msg);err!=nil{
// 			log.Println("error while publishing... %v",err)
// 		}
// 	}
// }
