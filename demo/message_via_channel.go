package main

import (
	"fmt"
	"github.com/realtime-chat-webapp-backend/models"
	"math/rand"
	"time"
)

// Demo for sending message object via channel
type MsgClient struct {
	msgChan chan models.Message
}

func NewMsgClient() *MsgClient {
	return &MsgClient{
		msgChan: make(chan models.Message),
	}
}

// Receive
func (msgClient *MsgClient) write() {
	for msg := range msgClient.msgChan {
		fmt.Printf("%#v\n", msg)
	}
}

// Send
func (msgClient *MsgClient) subscribeMessages() {
	for {
		time.Sleep(randomInSecond())
		msgClient.msgChan <- models.Message{"message add", ""}
	}
}

// Send
func (msgClient *MsgClient) subscribeChannels() {
	for {
		time.Sleep(randomInSecond())
		msgClient.msgChan <- models.Message{"channel add", ""}
	}
}

func randomInSecond() time.Duration {
	return time.Millisecond * time.Duration(rand.Intn(1000))
}

func SendMsgViaChannel() {
	msgClient := NewMsgClient()
	go msgClient.subscribeChannels() // send data to channel
	go msgClient.subscribeMessages() // send data to channel
	msgClient.write()                // receive data from channel
}
