package ffprobe

// Format information.
type Format struct {
	Filename       string      `json:"filename"`
	NBStreams      int         `json:"nb_streams"`
	NBPrograms     int         `json:"nb_programs"`
	FormatName     string      `json:"format_name"`
	FormatLongName string      `json:"format_long_name"`
	StartTime      string      `json:"start_time"`
	Duration       string      `json:"duration"`
	Size           string      `json:"size"`
	BitRate        string      `json:"bit_rate"`
	ProbeScore     int         `json:"probe_score"`
	Tags           *FormatTags `json:"tags"`
}

type FormatTags struct {
	MajorBrand       string `json:"major_brand"`
	MinorVersion     string `json:"minor_version"`
	CompatibleBrands string `json:"compatible_brands"`
	CreationTime     string `json:"creation_time"`
}
