package main

import (
	"fmt"
	"net/http"
)

func main() {

	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println("Array is ", arr)

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

	var value_data = helloWorldWithText("Alex")
	fmt.Println(value_data)

}

func reverseWords(sentence string) string {

	// Replace this placeholder return statement with your code
	words := make([]string, 0)
	currentWord := ""

	for _, char := range sentence {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		words = append(words, currentWord)
	}

	reversedSentence := ""
	for i := len(words) - 1; i >= 0; i-- {
		reversedSentence += words[i]
		if i != 0 {
			reversedSentence += " "
		}
	}

	return reversedSentence
}

func helloWorldWithText(helloText string) string {
	return "Hello my value : " + helloText
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

func sayHelloWithFormat(name string) string {
	return fmt.Sprintf("Hello %d ", name)
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
