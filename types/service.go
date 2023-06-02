package types

type ServicePayload struct {
	Name                   string `json:"name"`
	Description            string `json:"description"`
	AutoResolve            bool   `json:"auto_resolve_timeout"`
	AcknowledgementTimeout int    `json:"acknowledgement_timeout"`
}
