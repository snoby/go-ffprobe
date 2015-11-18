package ffprobe

const (
	Video string = "video"
	Audio        = "audio"
)

// Stream information.
type Stream struct {
	Index            int    `json:"index"`
	CodecName        string `json:"codec_name"`
	CodecLongName    string `json:"codec_long_name"`
	CodecType        string `json:"codec_type"`
	CodecTimeBase    string `json:"codec_time_base"`
	CodecTagString   string `json:"codec_tag_string"`
	CodecTag         string `json:"codec_tag"`
	RFrameRate       string `json:"r_frame_rate"`
	AvgFrameRate     string `json:"avg_frame_rate"`
	TimeBase         string `json:"time_base"`
	StartPts         int    `json:"start_pts"`
	StartTime        string `json:"start_time"`
	DurationTs       uint64 `json:"duration_ts"`
	Duration         string `json:"duration"`
	BitRate          string `json:"bit_rate"`
	BitsPerRawSample string `json:"bits_per_raw_sample"`
	NbFrames         string `json:"nb_frames"`

	Disposition StreamDisposition `json:"disposition,omitempty"`
	Tags        StreamTags        `json:"tags,omitempty"`

	Profile            string `json:"profile,omitempty"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	HasBFrames         int    `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
	PixFmt             string `json:"pix_fmt,omitempty"`
	Level              int    `json:"level,omitempty"`
	ColorRange         string `json:"color_range,omitempty"`
	ColorSpace         string `json:"color_space,omitempty"`

	SampleFmt     string `json:"sample_fmt,omitempty"`
	SampleRate    string `json:"sample_rate,omitempty"`
	Channels      int    `json:"channels,omitempty"`
	ChannelLayout string `json:"channel_layout,omitempty"`
	BitsPerSample int    `json:"bits_per_sample,omitempty"`
}

// Type returns the type of stream.
func (s *Stream) Type() string {
	if s.CodecType == "video" {
		return Video
	}
	if s.CodecType == "audio" {
		return Audio
	}

	return ""
}

type StreamDisposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
}

type StreamTags struct {
	CreationTime string `json:"creation_time,omitempty"`
	Language     string `json:"language,omitempty"`
	Encoder      string `json:"encoder,omitempty"`
}
