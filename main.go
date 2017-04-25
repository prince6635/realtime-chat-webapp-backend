package main

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/realtime-chat-webapp-backend/controllers"
	"net/http"
	"fmt"
	"log"
)

// For realtime chat web app
func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "realtimechat",
	})
	if err != nil {
		fmt.Printf("Error while connectiong to RethinkDB: %#v\n", err)
		log.Panic(err.Error()) // App couldn't run without connecting to the DB
	}

	router := controllers.NewRouter(session)
	router.Handle("channel add", controllers.AddChannel)

	http.Handle("/", router)
	// 8080 is used by RethinkDB management UI
	http.ListenAndServe(":4000", nil)
}

/* For demo purpose
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // disable same origion policy for now
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	// http server
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	// WebSocket server
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//readMessagesDemo(socket)
	// Use gorilla package to encode/decode JSON data
	readMessages(socket)
}

// To test: http://jsbin.com and choose "javascript"->"ES6/Babel"
func readMessages(socket *websocket.Conn) {
	for {
		var recMsg models.Message
		err := socket.ReadJSON(&recMsg)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%#v\n", recMsg)

		switch recMsg.Name {
		case "channel add":
			channel, _ := utils.AddChannel(recMsg.Data)

			// TODO: save to database
			// Send success-saved message to the client
			var sendMsg models.Message
			sendMsg.Name = "channel add"
			sendMsg.Data = channel
			err := socket.WriteJSON(sendMsg)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%#v\n", recMsg)
		case "channel subscribe":
			go utils.SubscribeChannel(socket)
		}
	}
}

// To test: http://websocket.org/echo.html
func readMessagesDemo(socket *websocket.Conn) {
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(msg))

		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
*/
