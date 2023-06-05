package types

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
