package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func DecodeEncode() {
	recRawMsg := []byte(`{"name":"channel add",` +
		`"data":{"name":"Hardware Support"}}`)
	// Decode received raw message
	var recMsg Message

	// gorilla package has ReadJSON and writeJSON to handle Marshall & Unmarshall functions.
	err := json.Unmarshal(recRawMsg, &recMsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", recMsg)

	// Extract Channel object
	if recMsg.Name == "channel add" {
		addChannel(recMsg.Data)
		channel, _ := addChannelWithDecodeLib(recMsg.Data)

		// TODO: save to database
		// Send success-saved message to the client
		var sendMsg Message
		sendMsg.Name = "channel add"
		sendMsg.Data = channel
		sendRawMsg, err := json.Marshal(sendMsg)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(sendRawMsg))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	channelMap := data.(map[string]interface{}) // .() to assert the type
	channel.Name = channelMap["name"].(string)
	channel.Id = "1"

	fmt.Printf("%#v\n", channel)
	return channel, nil
}

func addChannelWithDecodeLib(data interface{}) (Channel, error) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}

	channel.Id = "1"

	fmt.Printf("%#v\n", channel)
	return channel, nil
}
