package controllers

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/gorilla/websocket"
	"github.com/realtime-chat-webapp-backend/models"
)

type FindHandler func(string) (Handler, bool)

// Responsibilities: send & receive messages to Browser
type Client struct {
	msgChan     chan models.Message
	socket      *websocket.Conn
	findHandler FindHandler // Find the corresponding handler in Router
	session *r.Session
}

func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		msgChan:     make(chan models.Message),
		socket:      socket,
		findHandler: findHandler,
		session: session,
	}
}

func (client *Client) Write() {
	for msg := range client.msgChan {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}

	client.socket.Close() // avoid channel leak
}

func (client *Client) Read() {
	var msg models.Message
	for {
		if err := client.socket.ReadJSON(&msg); err != nil {
			break
		}

		// decide which function to call
		if handler, found := client.findHandler(msg.Name); found {
			handler(client, msg.Data)
		}
	}

	client.socket.Close()
}
