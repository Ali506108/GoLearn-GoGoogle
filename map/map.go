package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello hwo are you !")

	kafka := make(map[string]int)
	kafka["localhost:9092"] = 9092
	fmt.Println(kafka)

	country := make(map[string]int64)

	country["United Kingdem"] = 1_100
	country["United States of America"] = 350

	fmt.Println(country)
	fmt.Printf("\n country with name %d", country["United Kingdem"])

	delete(country, "United Kingdem")

	fmt.Println("map after deleting ", country)

	clear(country)

	fmt.Println(country)

	my_map := map[string]int{
		"Alex":      20,
		"Alan":      20,
		"Elizabeth": 22,
		"Mark":      24,
		"Jesica":    34,
	}

	for i, v := range my_map {
		fmt.Printf("My name is %d and i'm %d years old\n", i, v)
	}

	if my_map == nil {
		fmt.Println("Error")
	} else {
		fmt.Println("Okay")
	}

	myMap5 := make(map[string]map[string]string)

	new_map := make(map[string]string)
	new_map["Alex"] = "U.K"

	myMap5["Country"] = new_map

	fmt.Println(myMap5)

	messages := "Hello how are you"

	for i, v := range messages {
		fmt.Println(i, v)
		fmt.Printf("Index : %d , Rune: %c\n", i, v)

	}

}
