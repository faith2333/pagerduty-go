package pagerduty_go

import (
	"sync"
)

type pagerduty struct {
	lock       *sync.RWMutex
	token      string
	restClient IRESTClient
}

func NewPagerDuty(token string) Interface {
	return &pagerduty{
		token:      token,
		restClient: NewDefaultRestClient(),
	}
}

func (pd *pagerduty) WithRESTClient(restClient IRESTClient) Interface {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	pd.restClient = restClient
	return pd
}
