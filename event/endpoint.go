package event

import "github.com/faith2333/pagerduty-go/types"

const (
	EnvHostDefault string         = "https://events.pagerduty.com"
	EndpointEvent  types.Endpoint = "/v2/enqueue"
)
