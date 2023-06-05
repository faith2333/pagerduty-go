package event

import "time"

type SendEventReq struct {
	Payload Payload `json:"payload"`
	// The GUID of one of your Events API V2 integrations.
	// This is the "Integration Key" listed on the Events API V2 integration's detail page.
	RoutingKey string `json:"routing_key"`
	// The type of event. Can be trigger, acknowledge or resolve.
	//  Allowed values: trigger acknowledge resolve
	EventAction Action `json:"event_action"`
	// Identifies the alert to trigger, acknowledge, or resolve. Required unless the event_type is trigger.
	DedupKey string `json:"dedup_key"`
	// The name of the monitoring client that is triggering this event. (This field is only used for trigger events.)
	Client string `json:"client"`
	// Links to be shown on the alert and/or corresponding incident.
	Links []Link `json:"links"`
	// Images to be displayed on the alert and/or corresponding incident.
	Images []Image `json:"images"`
}

type SendEventResp struct {
	// Returns "success" if successful, or a short error message in case of a failure.
	Status string `json:"status"`
	// The key used to correlate triggers, acknowledges, and resolves for the same alert.
	DedupKey string `json:"dedup_key"`
	// A description of the problem, or "Event processed" if successful.
	Message string `json:"message"`
}

type Payload struct {
	// A brief text summary of the event, used to generate the summaries/titles of any associated alerts.
	Summary string `json:"summary"`
	// The time at which the emitting tool detected or generated the event.
	Timestamp time.Time `json:"timestamp"`
	// The perceived severity of the status the event is describing withrespect to the affected system.
	// Allowed values: critical warning error info
	Severity Severity `json:"severity"`
	// The unique location of the affected system, preferably a hostname or FQDN.
	Source string `json:"source"`
	// Component of the source machine that is responsible for the event.
	Component string `json:"component"`
	// Logical grouping of components of a service.
	Group string `json:"group"`
	// The class/type of the event.
	Class string `json:"class"`
	// Additional details about the event and affected system.
	CustomDetails map[string]interface{} `json:"custom_details"`
}
