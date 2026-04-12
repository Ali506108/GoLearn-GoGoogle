package main

import (
	"fmt"
)

func main() {

	country := make(map[string]int64)

	country["United Kingdem"] = 1_100
	country["United States of America"] = 350

	fmt.Println(country)
	fmt.Printf("\n country with name %d", country["United Kingdem"])

}
