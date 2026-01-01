package main

import (
	"fmt"
)

func main() {

	fmt.Print("Hellp how are youy ? : ")
	var answer string
	fmt.Scan(&answer)
	fmt.Println("you're doing ", answer)

	fmt.Print("Tell you're name: ")
	var answerForFunc string
	fmt.Scan(&answerForFunc)

	value := helloWorld(answerForFunc)
	fmt.Println("Value is ", value)

}

func helloWorld(value string) string {
	return value + " Hello from helloWorld class"
}
