package types

const (
	APIBase = "https://api.pagerduty.com"

	EndpointUsers              Endpoint = APIBase + "/users"
	EndpointServices           Endpoint = APIBase + "/services"
	EndpointTeams              Endpoint = APIBase + "/teams"
	EndpointSchedules          Endpoint = APIBase + "/schedules"
	EndpointIncidents          Endpoint = APIBase + "/incidents"
	EndpointEscalationPolicies Endpoint = APIBase + "/escalation_policies"
	EndpointLogEntries         Endpoint = APIBase + "/log_entries"
)

var AllEndpoints = []Endpoint{EndpointUsers, EndpointServices, EndpointTeams, EndpointSchedules, EndpointIncidents, EndpointEscalationPolicies, EndpointLogEntries}

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
