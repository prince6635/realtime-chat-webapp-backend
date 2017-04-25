package controllers

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/realtime-chat-webapp-backend/models"
)

func AddChannel(client *Client, data interface{}) {
	var channel models.Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		// So we could show the error in browser!
		client.msgChan <- models.Message{"error", err.Error()}
		return
	}
	fmt.Printf("Added channel: %#v\n", channel)

	// Insert into RethinkDB
	// DB operation should
	go func() {
		err = r.Table("channel").
				Insert(channel).
				Exec(client.session)
		if err != nil {
			client.msgChan <- models.Message{"error", err.Error()}
		}
	}()

	/*
	// for Demo: verify we could send data from handler to browser
	var msg models.Message
	msg.Name = "channel add"
	channel.Id = "ABC123"
	msg.Data = channel
	client.msgChan <- msg
	*/
}
