package ffprobe

import (
	"encoding/json"
	"os/exec"
)

type container struct {
	Streams []*Stream `json:"streams"`
}

// Handler wraps a file.
type Handler struct {
	File string
}

// Streams returns the stream information.
func (h *Handler) Streams() (streams []*Stream, err error) {
	var (
		rawData []byte
	)

	cmd := exec.Command(
		"ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		h.File,
	)

	rawData, err = cmd.Output()
	if err != nil {
		return
	}

	c := new(container)
	if err = json.Unmarshal(rawData, &c); err != nil {
		return
	}

	streams = c.Streams

	return
}
