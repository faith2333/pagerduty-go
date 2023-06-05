package pagerduty_go

import (
	"context"
	"github.com/faith2333/pagerduty-go/event"
)

func (pd *pagerduty) SendEvent(ctx context.Context, payload *event.SendEventReq) (*event.SendEventResp, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithHost(event.EnvHostDefault).
		WithEndpoint(event.EndpointEvent).
		WithBody(payload).
		POST().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	return pd.respToSendEventResp(resp)
}

func (pd *pagerduty) respToSendEventResp(resp []byte) (sendResp *event.SendEventResp, err error) {
	err = pd.transformJson(resp, &sendResp)
	return
}
