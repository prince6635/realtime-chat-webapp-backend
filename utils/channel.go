package utils

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/realtime-chat-webapp-backend/models"
	"time"
)

func AddChannel(data interface{}) (models.Channel, error) {
	var channel models.Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}

	channel.Id = "1"

	fmt.Printf("Added channel: %#v\n", channel)
	return channel, nil
}

/* !!! FLOW:
1, client connects the server and send "channel subscribe" message
2, server gets "channel subscribe" message, call SubscribeChannel to make sure
	whenever there's a change in its subscribed channel ist happened, it'll notify the client
3, server sends "channel add" to the client to tell that its subscription is successful.
*/
func SubscribeChannel(socket *websocket.Conn) {
	fmt.Printf("Subscribed channel by: %#v\n", socket)

	// TODO: query rethinkDB with the feature: changefeed,
	// it'll look up initial channels, then keep
	// blocking and waiting for channel changes such as ADD, REMOVE, or EDIT
	for {
		time.Sleep(time.Second * 1)

		msg := models.Message{
			"channel add",
			models.Channel{"1", "Software Support"}}
		socket.WriteJSON(msg)
		fmt.Println("sent newly added channel.")
	}
}
