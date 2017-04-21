package controllers

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/realtime-chat-webapp-backend/models"
)

func AddChannel(client *Client, data interface{}) {
	var channel models.Channel
	mapstructure.Decode(data, &channel)
	fmt.Printf("Added channel: %#v\n", channel)

	// TODO: insert into RethinkDB
	var msg models.Message
	msg.Name = "channel add"
	channel.Id = "ABC123"
	msg.Data = channel
	client.msgChan <- msg
}
