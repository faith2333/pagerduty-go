package types

type ScheduleAction struct {
	Type      Type   `json:"type"`
	At        At     `json:"at"`
	ToUrgency string `json:"to_urgency"`
}

type At struct {
	// Must be set to named_time.
	Type Type `json:"type"`
	// Designates either the start or the end of support hours.
	// Allowed values: support_hours_start support_hours_end
	Name string `json:"name"`
}
