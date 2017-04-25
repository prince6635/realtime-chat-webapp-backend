package main

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"fmt"
	"time"
)

// ===Still keep listening after printing "Browser closes..."===
func subscribeChangefeedForCannot(session *r.Session) {
	var change r.ChangeResponse
	cursor, _ := r.Table("channel").
		Changes().
		Run(session)

	for cursor.Next(&change) {
		// In actual app, send update to client
		fmt.Printf("%#v\n", change.NewValue)
	}
}

func cannotStopBlockingGoroutine(session *r.Session)  {
	go subscribeChangefeedForCannot(session)
	// Sleep to keep app running
	time.Sleep(time.Second * 5)
	fmt.Println("Browser closes... websocket closes...")
	time.Sleep(time.Second * 10000)
}
// ==============================================================

// ===Stop listening changefeed after printing "Browser closes..."===
func subscribeChangefeedForCan(session *r.Session, stop <-chan bool) { // <-chan means it's a receive-only channel
	changeResponse := make(chan r.ChangeResponse)
	cursor, _ := r.Table("channel").
			Changes().
			Run(session)

	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			// In actual app, send update to client
			//fmt.Printf("%#v\n", change.NewValue)
			changeResponse <- change
		}
		fmt.Println("Exiting cursor goroutine...")
	}()

	for {
		select {
		case change := <-changeResponse:
			fmt.Printf("%#v\n", change.NewValue)
		case <-stop:
			fmt.Println("Closing cursor...")
			cursor.Close()
			return
		}
	}
}

func canStopBlockingGoroutine(session *r.Session) {
	stop := make(chan bool)

	go subscribeChangefeedForCan(session, stop)
	// Sleep to keep app running
	time.Sleep(time.Second * 5)
	fmt.Println("Sending stop signal")
	stop <- true
	fmt.Println("Browser closes... websocket closes...")
	time.Sleep(time.Second * 10000)
}
// ==============================================================

func StopBlockingGoroutine() {
	session, _ := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: "realtimechat",
	})

	//cannotStopBlockingGoroutine(session)
	canStopBlockingGoroutine(session)
}