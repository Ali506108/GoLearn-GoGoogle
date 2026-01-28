package main

import (
	"encoding/json"
	"log"
	"reflect"
	"time"
)

type User struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Age       int       `json:"age"`
}

type Profile struct {
	City     string  `json:"city"`
	Temp     float64 `json:"temp"`
	UserInfo User    `json:"user"`
}

func newProfile() Profile {
	return Profile{
		City: "Boston",
		Temp: 38.4,
		UserInfo: User{
			Username:  "Alex-uk",
			Email:     "aliduisen77@gmail.com",
			CreatedAt: time.Now().UTC(),
			Age:       20,
		},
	}
}

func main() {
	profile := newProfile()
	jsonData, err := json.MarshalIndent(profile, "", " ")
	if err != nil {
		log.Fatalf("error marshaling record: %v", err)
	}

	log.Printf("Profile created : %s\n", jsonData)

	t := reflect.TypeOf(profile)
	v := reflect.ValueOf(profile)

	log.Println("Inspect struct fields ")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		log.Printf("   %s (%s) = %v\n", field.Name, field.Type, value.Interface())
	}

}
