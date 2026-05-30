package main

import (
	"fmt"
)

type Avatar struct {
	URL  string
	Size int64
}

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  Avatar
}

func updateUrl(client *Client) {
	client.IMG.URL = "https://golang.com"
}

func updateSize(client *Client) {
	client.IMG.Size = 256
}

func main() {

	client := Client{
		ID:   1,
		Name: "Alex",
		Age:  20,
		IMG: Avatar{
			URL:  "",
			Size: 1,
		},
	}

	fmt.Printf("Data client: %+v\n", client)

	updateUrl(&client)
	fmt.Printf("With client update url : %+v\n", client)
	updateSize(&client)
	fmt.Printf("With client update size : %+v\n", client)

	seq := adder()
	fmt.Println(seq())
	fmt.Println(seq())

	fmt.Println(seq())

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i was : ", i)

	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}

}
