package types

type Notification struct {
}

type AutoPauseNotificationsParameters struct {
	// Indicates whether alerts should be automatically suspended when identified as transient
	// Default: false
	Enabled bool `json:"enabled"`
	// Indicates in seconds how long alerts should be suspended before triggering
	//
	//Allowed values: 120 180 300 600 900
	Timeout int `json:"timeout"`
}
