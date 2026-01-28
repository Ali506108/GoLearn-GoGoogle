package main

import (
	"fmt"
	"time"
)

type User struct {
	Username string
	Pasword  string
	email    string
	age      int
	data     time.Time
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 User
}

func main() {
	Data := Record{
		"London",
		-12.3465,
		User{
			"Alex-uk",
			"root_root",
			"aliduisen77@gmail.com",
			20,
			time.Now().UTC(),
		},
	}

	fmt.Printf("User %s created at %s\n", Data.Field3.Username,
		Data.Field3.data.Format("2006-01-02 15:04:05"),
	)
}
