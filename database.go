// This is a fake database to speed up prototyping
// with aim to simulate database interactions

package main

type UserProfile struct {
	Id       string
	Email    string
	Username string
	Token    string
}

// Creates an array storing UserProfile struct instances
// to simulate database

var database = map[string]UserProfile{
	"USER0": {
		Id:       "USER0",
		Email:    "example@domain.com",
		Username: "X_Sample",
		Token:    "123",
	},
	"USER1": {
		Id:       "USER1",
		Email:    "second.example@domain.com",
		Username: ".",
		Token:    "456",
	},
}
