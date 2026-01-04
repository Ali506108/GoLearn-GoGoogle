package main

import (
	"fmt"
	"time"
)

type User struct {
	Name     string
	LastName string
	Age      int
	isOnline bool
}

type FilmMedia struct {
	Title    string
	URL      string
	Language Language
	Data     time.Time
}

type Language struct {
	name string
}

type Movie struct {
	Title string
}

func newMediaFile(title, url, langName string, data time.Time) FilmMedia {
	return FilmMedia{
		Title: title,
		URL:   url,
		Language: Language{
			name: langName,
		},
		Data: data,
	}
}

func (f FilmMedia) printInfo() {
	fmt.Println("Data was : ", f.Title, f.Language.name)
}

type Media interface {
	Play()
}

func (m Movie) Play() {
	fmt.Println("Playing movie is ", m.Title)
}
func main() {
	const MAXRETRIES = 5

	var employeeID = 2
	fmt.Println(employeeID)
}
