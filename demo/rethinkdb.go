package main

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"fmt"
	"github.com/realtime-chat-webapp-backend/models"
)

func RethinkDB() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015", // !!!NOTE: it's an url, no space between localhost and 28015
		Database: "realtimechat",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Insert
	user := models.User{
		Name: "anonymous",
	}
	response, err := r.Table("user").
			Insert(user).
			RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Insert: %#v\n", response)
	newId := response.GeneratedKeys[0]
	fmt.Printf("Insert: %#v\n", newId)

	// Update
	user.Name = "Zi"
	response, err = r.Table("user").
		Get(newId).
		Update(user).
		RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Update: %#v\n", response)

	// Delete
	response, err = r.Table("user").
		Get(newId).
		Delete().
		RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Delete: %#v\n", response)

	// Changefeed (keep showing in console if you have any change in user table via http://localhost:8080/#dataexplorer)
	cursor, _ := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		fmt.Printf("%#v\n", changeResponse)
	}
}
