package models

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}
