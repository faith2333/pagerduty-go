package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

type Incident struct {
	BaseObject
	// The number of the incident. This is unique across your account.
	IncidentNumber int `json:"incident_number"`
	// The time the incident was first triggered.
	//   Example: 2019-12-01T20:00:00Z
	CreateAt time.Time `json:"create_at"`
	// The current status of the incident.
	//  Allowed values: triggered acknowledged resolved
	Status IncidentStatus `json:"status"`
	// A succinct description of the nature, symptoms, cause, or effect of the incident.
	Title string `json:"title"`
	// The list of pending_actions on the incident.
	// A pending_action object contains a type of action which can be one of escalate, unacknowledge, resolve or urgency_change.
	// A pending_action object contains at, the time at which the action will take place.
	// An urgency_change pending_action will contain to, the urgency that the incident will change to.
	PendingActions []PendingAction `json:"pending_actions"`
	// The incident's de-duplication key.
	IncidentKey string              `json:"incident_key"`
	Service     BaseObjectReference `json:"service"`
	// List of all assignments for this incident. This list will be empty if the Incident.status is resolved.
	Assignments []Assignment `json:"assignments"`
	// How the current incident assignments were decided. Note that direct_assignment incidents will not escalate up the attached escalation_policy
	//   Allowed values: escalation_policy  direct_assignment
	AssignedVia string `json:"assigned_via"`
	// List of all acknowledgements for this incident. This list will be empty if the Incident.status is resolved or triggered.
	Acknowledgements []Acknowledgement `json:"acknowledgements"`
	// The time the status of the incident last changed. If the incident is not currently acknowledged or resolved, this will be the incident's updated_at.
	//   Example: 2019-12-01T21:01:00Z
	LastStatusChangeAt time.Time `json:"last_status_change_at"`
	// The agent (user, service or integration) that created or modified the Incident Log Entry.
	LastStatusChangeBy   BaseObjectReference   `json:"last_status_change_by"`
	FirstTriggerLogEntry BaseObjectReference   `json:"first_trigger_log_entry"`
	EscalationPolicy     BaseObjectReference   `json:"escalation_policy"`
	Teams                []BaseObjectReference `json:"teams"`
	Priority             BaseObjectReference   `json:"priority"`
	// The current urgency of the incident.
	//   Allowed values: high low
	Urgency           Urgency            `json:"urgency"`
	ResolveReason     ResolveReason      `json:"resolve_reason"`
	AlertCounts       AlertCounts        `json:"alert_counts"`
	ConferenceBridge  ConferenceBridge   `json:"conference_bridge"`
	IncidentBody      IncidentBody       `json:"incident_body"`
	IncidentResponder IncidentResponder  `json:"incident_responder"`
	ResponderRequests []ResponderRequest `json:"responder_requests"`
	ResolvedAt        time.Time          `json:"resolved_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
}

const (
	IncidentStatusTriggered    IncidentStatus = "triggered"
	IncidentStatusAcknowledged IncidentStatus = "acknowledged"
	IncidentStatusResolved     IncidentStatus = "resolved"
)

type IncidentStatus string

func (is IncidentStatus) String() string {
	return string(is)
}

type IncidentUrgencyRule struct {
	Type string `json:"type"`
	// The incidents' urgency, if type is constant.
	Urgency             Urgency       `json:"urgency"`
	DuringSupportHours  []UrgencyRule `json:"during_support_hours"`
	OutsideSupportHours []UrgencyRule `json:"outside_support_hours"`
}

type UrgencyRule struct {
	Type    string  `json:"type"`
	Urgency Urgency `json:"urgency"`
}

const (
	UrgencyLow           Urgency = "low"
	UrgencyHigh          Urgency = "high"
	UrgencySeverityBased Urgency = "severity_based"
)

type Urgency string

func (u Urgency) String() string {
	return string(u)
}

const (
	PendingActionTypeUnacknowledged PendingActionType = "unacknowlege"
	PendingActionTypeEscalate       PendingActionType = "escalate"
	PendingActionTypeResolve        PendingActionType = "resolve"
	PendingUrgencyChange            PendingActionType = "urgency_change"
)

type PendingActionType string

func (pat PendingActionType) String() string {
	return string(pat)
}

type PendingAction struct {
	Type PendingActionType `json:"type"`
	At   time.Time         `json:"at"`
}

type ResolveReason struct {
	// The reason the incident was resolved. The only reason currently supported is merge.
	//   Allowed value: merge_resolve_reason
	//   Default: merge_resolve_reason
	Type     string              `json:"type"`
	Incident BaseObjectReference `json:"incident"`
}

type ConferenceBridge struct {
	// The phone number of the conference call for the conference bridge.
	// Phone numbers should be formatted like +1 415-555-1212,,,,1234#,
	// where a comma (,) represents a one-second wait and pound (#) completes access code input.
	ConferenceNumber string `json:"conference_number"`
	// A URL for the conference bridge. This could be a link to a web conference or Slack channel.
	ConferenceURL string `json:"conference_url"`
}

type IncidentBody struct {
	// Allowed value: incident_body
	Type string `json:"type"`
	// Additional incident details.
	Details string `json:"details"`
}

type IncidentResponder struct {
	// The status of the responder being added to the incident
	//  Example: pending
	State     string              `json:"state"`
	User      BaseObjectReference `json:"user"`
	Incident  BaseObjectReference `json:"incident"`
	UpdatedAt time.Time           `json:"updated_at"`
	Message   string              `json:"message"`
	Requester BaseObjectReference `json:"requester"`
	RequestAt time.Time           `json:"requestAt"`
}

type ResponderRequest struct {
	Incident    BaseObjectReference `json:"incident"`
	Requester   BaseObjectReference `json:"requester"`
	RequestedAt string              `json:"requested_at"`
	Message     string              `json:"message"`
}

type ResponderRequestTarget struct {
	Type               string              `json:"type"`
	ID                 string              `json:"id"`
	Summary            string              `json:"summary"`
	IncidentResponders []IncidentResponder `json:"incident_responders"`
}

type ListIncidentsReq struct {
	// The number of results per page.
	Limit int `json:"limit"`
	// Offset to start pagination search results.
	Offset int `json:"offset"`
	// default the total field in pagination responses is set to null to provide the fastest possible response times.
	//  Set total to true for this field to be populated.
	//  Default: false
	Total bool `json:"total"`
	// When set to all, the Since and Until parameters and Defaults are ignored.
	//  Allowed value: all
	DataRange string `json:"data_range"`
	// Incident de-duplication key. Incidents with child alerts do not have an incident key;
	// querying by incident key will return incidents whose alerts have alert_key matching the given incident key.
	IncidentKey string `json:"incident_key"`
	// Array of additional details to include.
	//  Allowed values: users services first_trigger_log_entries escalation_policies teams assignees acknowledgers priorities conference_bridge
	Include string `json:"include[]"`
	// Returns only the incidents associated with the passed service(s). This expects one or more service IDs.
	ServiceIds string `json:"service_ids[]"`
	// The start of the date range over which you want to search. Maximum range is 6 months and default is 1 month.
	Since string `json:"since"`
	// Used to specify both the field you wish to sort the results on (incident_number/created_at/resolved_at/urgency),
	// as well as the direction (asc/desc) of the results. The sort_by field and direction should be separated by a colon.
	// A maximum of two fields can be included, separated by a comma. Sort direction defaults to ascending.
	// The account must have the urgencies ability to sort by the urgency.
	//  <= 2 items
	SortBy []string `json:"sort_by"`
	// Return only incidents with the given statuses. To query multiple statuses, pass statuses[] more than once,
	// for example: https://api.pagerduty.com/incidents?statuses[]=triggered&statuses[]=acknowledged.
	// (More status codes may be introduced in the future.)
	//  Allowed values: triggered acknowledged resolved
	Statuses string `json:"statuses[]"`
	// An array of team IDs. Only results related to these teams will be returned.
	// Account must have the teams ability to use this parameter.
	TeamIds []string `json:"team_ids[]"`
	// Time zone in which results will be rendered. This will default to the account time zone.
	TimeZone string `json:"timeZone"`
	// The end of the date range over which you want to search. Maximum range is 6 months and default is 1 month.
	Until string `json:"until"`
	// Array of the urgencies of the incidents to be returned. Defaults to all urgencies. Account must have the urgencies ability to do this.
	// Allowed values: high low
	Urgencies string `json:"urgencies[]"`
	// Returns only the incidents currently assigned to the passed user(s).
	// This expects one or more user IDs. Note: When using the assigned_to_user filter,
	// you will only receive incidents with statuses of triggered or acknowledged.
	// This is because resolved incidents are not assigned to any user.
	UserIds []string `json:"user_ids[]"`
}

func (req *ListIncidentsReq) AsMap() (map[string]interface{}, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "marshal request failed")
	}

	reqMap := make(map[string]interface{})
	err = json.Unmarshal(reqBytes, &reqMap)
	if err != nil {
		return nil, err
	}
	return reqMap, nil
}

type ListIncidentsResp struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	// Indicates if there are additional records to return
	More bool `json:"more"`
	// The total number of records matching the given query.
	Total     int        `json:"total"`
	Incidents []Incident `json:"incidents"`
}

type CreateIncidentPayload struct {
	// Allowed value: incident
	Type Type `json:"type"`
	// A succinct description of the nature, symptoms, cause, or effect of the incident.
	Title    string              `json:"title"`
	Service  BaseObjectReference `json:"service"`
	Priority BaseObjectReference `json:"priority"`
	// The urgency of the incident
	//  Allowed values: high low
	Urgency          Urgency             `json:"urgency"`
	Body             IncidentBody        `json:"body"`
	IncidentKey      string              `json:"incident_key"`
	Assignments      []Assignment        `json:"assignments"`
	EscalationPolicy BaseObjectReference `json:"escalation_policy"`
	ConferenceBridge BaseObjectReference `json:"conference_bridge"`
}

type GetIncidentPayload struct {
	Incident *Incident `json:"incident"`
}
