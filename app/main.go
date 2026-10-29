package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"notify-send",
		"-u", "critical",
		"-a", "Watcher",
		"Critical notification!",
		"this is the body of the notification",
	)

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
