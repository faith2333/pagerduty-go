package pagerduty_go

import "net/http"

type defaultPagerDutyClient struct {
	request    *http.Request
	httpClient *http.Client
}

func NewDefaultPagerDutyClient() Interface {
	return &defaultPagerDutyClient{
		httpClient: &http.Client{},
	}
}

func (dClient *defaultPagerDutyClient) WithHost(host string) Interface {
	return dClient
}
