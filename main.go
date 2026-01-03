package main

import (
	"fmt"
	"net/http"
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

	numberFromFunction := addNumber(4, 7)
	fmt.Println(numberFromFunction)

	netOperation()

	fmt.Println()

	fmt.Print("give me id: ")
	var id int
	fmt.Scan(&id)
	getData(id)

	var data = sayHello("Alex")
	fmt.Println(data)

}

func helloWorld(value string) string {
	return value + " Hello from helloWorld class"
}

func addNumber(valueOne, valueTwo int) int {
	return valueOne + valueTwo
}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %d", name)
}

func getData(id int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
	defer resp.Body.Close()
	fmt.Println("Response statuse : ", resp.Status, " ", resp.StatusCode, " ", resp.Request.URL)
}

func netOperation() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	defer resp.Body.Close()

	fmt.Println("Response statuse : ", resp.Status)
}
