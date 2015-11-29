package ffprobe

import (
	"encoding/json"
	"errors"
	"os/exec"
)

type container struct {
	Streams []*Stream `json:"streams"`
	Format  *Format   `json:"format"`
}

// Handler wraps a file.
type Handler struct {
	File string
	data *container
}

func (h *Handler) getData() (err error) {
	if h.data != nil && h.data.Streams != nil && h.data.Format != nil {
		return
	}

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

	h.data = c

	return
}

// Streams returns the stream information.
func (h *Handler) Streams() (streams []*Stream, err error) {
	if err = h.getData(); err != nil {
		return
	}

	if h.data.Streams == nil {
		err = errors.New("could not get stream info")
		return
	}

	streams = h.data.Streams

	return
}

// Format returns the format information.
func (h *Handler) Format() (format *Format, err error) {
	if err = h.getData(); err != nil {
		return
	}

	if h.data.Format == nil {
		err = errors.New("could not get format info")
		return
	}

	format = h.data.Format

	return
}
