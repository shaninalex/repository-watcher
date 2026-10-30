package notifier

import (
	"os/exec"
)

func SystemNotify(level, title, text string) error {
	cmd := exec.Command(
		"notify-send",
		"-u", level,
		"-a", "Watcher",
		title,
		text,
	)
	return cmd.Run()
}
