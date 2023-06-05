package types

type BaseObject struct {
	ID   string `json:"id"`
	Type Type   `json:"type"`
	// A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client.
	// In many cases, this will be identical to name, though it is not intended to be an identifier.
	Summary string `json:"summary"`
	// the API show URL at which the object is accessible
	Self string `json:"self"`
	// a URL at which the entity is uniquely displayed in the Web app
	HTMLUrl string `json:"html_url"`
}

// BaseObjectReference the reference object
type BaseObjectReference struct {
	BaseObject
	Type TypeReference `json:"type"`
}

const (
	TypeUser    Type = "user"
	TypeService Type = "service"

	TypeLicense          Type = "license"
	TypeEscalationPolicy Type = "escalation_policy"
	TypeNotificationRule Type = "notification_rule"
	TypeResponsePlay     Type = "response_play"
	TypeUrgencyChange    Type = "urgency_change"
)

type Type string

func (t Type) String() string {
	return string(t)
}

func (t Type) Reference() TypeReference {
	return TypeReference(t.String() + "_reference")
}

type TypeReference string

func (t TypeReference) String() string {
	return string(t)
}
