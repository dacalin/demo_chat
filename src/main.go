package main

import (
	"context"
	"github.com/dacalin/demo_chat/bootstrap"
	"github.com/dacalin/ws_gateway"
	_connection_id "github.com/dacalin/ws_gateway/models/connection_id"
	"strconv"
)

func main() {
	ctx := context.Background()
	config := bootstrap.GetConfig()
	redidAddress := config.RedisHost + ":" + strconv.Itoa(config.RedisPort)

	wsGatewayConnection := ws_gateway.CreateConnectionGateway()
	wsServer := ws_gateway.CreateServer(redidAddress, config.WsPingIntervalSeconds, ctx)

	wsServer.OnConnect(func(connectionId _connection_id.ConnectionId, params map[string]string) {
		wsGatewayConnection.SetGroup(connectionId, "demo-room")
	})

	wsServer.OnMessage(
		func(connectionId _connection_id.ConnectionId, data []byte) {
			wsGatewayConnection.Broadcast("demo-room", data)
		})

	wsServer.Run(config.WsPort)

}
