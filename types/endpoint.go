package types

const (
	HostDefault = "https://api.pagerduty.com"

	EndpointUsers              Endpoint = "/users"
	EndpointServices           Endpoint = "/services"
	EndpointTeams              Endpoint = "/teams"
	EndpointSchedules          Endpoint = "/schedules"
	EndpointIncidents          Endpoint = "/incidents"
	EndpointEscalationPolicies Endpoint = "/escalation_policies"
	EndpointLogEntries         Endpoint = "/log_entries"
)

var AllEndpoints = []Endpoint{EndpointUsers, EndpointServices, EndpointTeams,
	EndpointSchedules, EndpointIncidents, EndpointEscalationPolicies, EndpointLogEntries}

type Endpoint string

func (e Endpoint) String() string {
	return string(e)
}

func (e Endpoint) IsUsers() bool {
	return e == EndpointUsers
}

func (e Endpoint) IsServices() bool {
	return e == EndpointServices
}

func (e Endpoint) IsTeams() bool {
	return e == EndpointTeams
}

func (e Endpoint) IsSchedules() bool {
	return e == EndpointSchedules
}

func (e Endpoint) IsIncidents() bool {
	return e == EndpointIncidents
}

func (e Endpoint) IsEscalationPolices() bool {
	return e == EndpointEscalationPolicies
}

func (e Endpoint) IsLogEntries() bool {
	return e == EndpointLogEntries
}

func (e Endpoint) IsSupported() bool {
	for _, endpointSupported := range AllEndpoints {
		if endpointSupported == e {
			return true
		}
	}
	return false
}
