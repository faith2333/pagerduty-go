package types

type ContactMethod struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Address string `json:"address"`
}
