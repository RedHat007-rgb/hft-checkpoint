package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/RedHat007-rgb/hft-checkpoint/packages/golib/redis"
	"github.com/RedHat007-rgb/hft-checkpoint/packages/golib/ws"
	pb "github.com/RedHat007-rgb/hft-checkpoint/packages/proto/proto/ticker"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)


type TickerServer struct{
	pb.UnimplementedTickerServiceServer
}

func (s *TickerServer) Subscribe(ctx context.Context, req *pb.TickerRequest) (*pb.TickerAck, error) {
    channel := "binance." + req.Symbol + ".ticker"
    if err := redisClient.SetUser(ctx, channel, 1); err != nil {
        log.Println("error while setting:", err)
        return nil, err
    }
    return &pb.TickerAck{
        Symbol: req.Symbol,
    }, nil
}


var redisClient = redis.NewConnection()


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
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil{
		log.Println("error while creating a server %v",err)
	}
	grpc:=grpc.NewServer()

	ctx := context.Background()
	tokens := []string{"btcusdt", "solusdt", "ethusdt"}
	BinanceWsConnections(tokens)

	for index := range binanceWsConns {
		wsConn := binanceWsConns[index].conn
		token := tokens[index]
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
	
	pb.RegisterTickerServiceServer(grpc, &TickerServer{})
	log.Println("listeniong on port :50051")
	grpc.Serve(lis)
}



