package ffprobe

import (
	"os/exec"
)

// IsAvailable returns true if ffprobe command is available;
// false otherwise.
func IsAvailable() bool {
	cmd := exec.Command("ffprobe", "-version")
	return cmd.Run() == nil
}

// File returns a handler for a file.
func File(file string) (handler *Handler) {
	return &Handler{
		File: file,
	}
}
