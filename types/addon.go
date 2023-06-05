package types

type Addon struct {
	BaseObject
	// The URL source of the Addon
	Src string `json:"src"`
	// The user entered name of the Addon.
	Name string `json:"name"`
}
