package main

import (
	"context"
	"github.com/dacalin/demo_chat/bootstrap"
	"github.com/dacalin/ws_gateway"
	_connection_id "github.com/dacalin/ws_gateway/models/connection_id"
)

func main() {
	ctx := context.Background()
	config := bootstrap.GetConfig()

	WSConfig := ws_gateway.Config{
		Driver:         "gws",
		EnableDebugLog: true,
		GWSDriver: ws_gateway.GWSDriverConfig{
			RedisHost:           config.RedisHost,
			RedisPort:           config.RedisPort,
			PingIntervalSeconds: config.WsPingIntervalSeconds,
			WSRoute:             "connect",
		},
	}

	wsServer, wsGatewayConnection, err := ws_gateway.Create(WSConfig, ctx)
	if err != nil {
		panic(err)
	}

	wsServer.OnConnect(func(connectionId _connection_id.ConnectionId, params map[string]string) {
		wsGatewayConnection.SetGroup(connectionId, "demo-room")
	})

	wsServer.OnMessage(
		func(connectionId _connection_id.ConnectionId, data []byte) {
			wsGatewayConnection.Broadcast("demo-room", data)
		})

	wsServer.Run(config.WsPort)

}
