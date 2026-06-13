package main

import (
	"fmt"
	os_go "gogoogle/os"
	srv "gogoogle/os/server"
)

func main() {

	runApp := srv.ServerShow()
	main_d := os_go.SendMessage("Java is better")

	fmt.Println("Main_d : ", main_d)

	fmt.Println("RunApp : ", runApp)
}
