package types

const (
	SoundTypeHigh SoundType = "alert_high_urgency"
	SoundTypeLow  SoundType = "alert_low_urgency"
)

type SoundType string

func (st SoundType) String() string {
	return string(st)
}

type Sound struct {
	Type SoundType `json:"type"`
	// The sound file name.
	File string `json:"file"`
}
