package types

type Base struct {
	ID string `json:"id"`
	// A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client.
	// In many cases, this will be identical to name, though it is not intended to be an identifier.
	Summary string `json:"summary"`
	// the API show URL at which the object is accessible
	Self string `json:"self"`
	// a URL at which the entity is uniquely displayed in the Web app
	HTMLUrl string `json:"html_url"`
}
