package os

import (
	"log"
	"os/exec"
)

func SendMessage(msg string) string {
	return msg
}

func RunSimpleApp() {
	cmd := exec.Command("Chrome")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

}
