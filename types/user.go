package types

type UserPayload struct {
	Type           string          `json:"type"`
	Name           string          `json:"name"`
	Email          string          `json:"email"`
	Timezone       string          `json:"timezone"`
	Role           string          `json:"role"`
	ContactMethods []ContactMethod `json:"contact_methods,omitempty"`
	TeamIDs        []string        `json:"team_ids,omitempty"`
}
