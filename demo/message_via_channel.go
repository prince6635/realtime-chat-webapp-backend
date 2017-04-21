package main

import (
	"github.com/realtime-chat-webapp-backend/models"
	"time"
	"math/rand"
	"fmt"
)

// Demo for sending message object via channel
type MsgClient struct {
	msgChann chan models.Message
}

func NewMsgClient() *MsgClient {
	return &MsgClient{
		msgChann: make(chan models.Message),
	}
}

// Receive
func (msgClient *MsgClient) write() {
	for msg := range msgClient.msgChann {
		fmt.Printf("%#v\n", msg)
	}
}

// Send
func (msgClient *MsgClient) subscribeMessages() {
	for {
		time.Sleep(randomInSecond())
		msgClient.msgChann <- models.Message{"message add", ""}
	}
}

// Send
func (msgClient *MsgClient) subscribeChannels() {
	for {
		time.Sleep(randomInSecond())
		msgClient.msgChann <- models.Message{"channel add", ""}
	}
}

func randomInSecond() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func SendMsgViaChannel() {
	msgClient := NewMsgClient()
	go msgClient.subscribeChannels() // send data to channel
	go msgClient.subscribeMessages() // send data to channel
	msgClient.write() // receive data from channel
}
