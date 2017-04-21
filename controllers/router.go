package controllers

import (
	"net/http"
	"github.com/gorilla/websocket"
)

type Handler func(*Client, interface{})

type Router struct {
	rules map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // disable same origion policy for now
	},
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// WebSocket server
	//socket, err := upgrader.Upgrade(w, req, nil)
}