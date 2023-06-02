package pagerduty_go

type pagerduty struct {
	restClient IRESTClient
}

func NewPagerDuty(restClient IRESTClient) Interface {
	return &pagerduty{
		restClient: restClient,
	}
}
