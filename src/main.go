package main

import (
	"context"
	"github.com/dacalin/demo_chat/bootstrap"
	"github.com/dacalin/ws_gateway"
	_connection_id "github.com/dacalin/ws_gateway/models/connection_id"
	"time"
)

func main() {
	ctx := context.Background()
	config := bootstrap.GetConfig()

	WSConfig := ws_gateway.Config{
		Driver:         ws_gateway.DRIVER_WS_GWS,
		EnableDebugLog: false,
		GWSDriver: ws_gateway.GWSDriverConfig{
			PubSub: ws_gateway.PubSubDriverConfig{
				Driver: ws_gateway.DRIVER_PUBSUB_REDIS,
				Host:   config.RedisHost,
				Port:   config.RedisPort,
			},
			PingIntervalSeconds: config.WsPingIntervalSeconds,
			WSRoute:             "connect",
		},
	}

	wsServer1, wsGatewayConnection1, err := ws_gateway.Create(WSConfig, ctx)
	if err != nil {
		panic(err)
	}

	wsServer1.OnConnect(func(connectionId _connection_id.ConnectionId, params map[string]string) {
		wsGatewayConnection1.SetGroup(connectionId, "demo-room")
	})

	wsServer1.OnMessage(
		func(connectionId _connection_id.ConnectionId, data []byte) {
			time.Sleep(10 * time.Second)
			wsGatewayConnection1.Broadcast("demo-room", data)
		})

	wsServer2, wsGatewayConnection2, err := ws_gateway.Create(WSConfig, ctx)
	if err != nil {
		panic(err)
	}
	wsServer2.OnConnect(func(connectionId _connection_id.ConnectionId, params map[string]string) {
		wsGatewayConnection2.SetGroup(connectionId, "demo-room")
	})

	wsServer2.OnMessage(
		func(connectionId _connection_id.ConnectionId, data []byte) {
			wsGatewayConnection2.Broadcast("demo-room", data)
		})

	go wsServer1.Run(config.WsPort)
	wsServer2.Run(config.WsPort + 1)

}
