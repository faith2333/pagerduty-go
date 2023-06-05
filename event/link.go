package event

type Link struct {
	// The link being attached to an incident or alert.
	Href string `json:"href"`
	// Optional information pertaining to this context link.
	Text string `json:"text"`
}
