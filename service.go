package pagerduty_go

import (
	"context"
	"encoding/json"
	"github.com/faith2333/pagerduty-go/types"
)

func (pd *pagerduty) GetService(ctx context.Context, id string) (*types.Service, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointServices).
		AddPath(id).
		GET().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetServiceResp(resp)
	if err != nil {
		return nil, err
	}

	return getResp.Service, nil
}

func (pd *pagerduty) CreateService(ctx context.Context, payload *types.CreateAndUpdateServicePayload) (*types.Service, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointServices).
		WithBody(payload).
		POST().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetServiceResp(resp)
	if err != nil {
		return nil, err
	}

	return getResp.Service, nil
}

func (pd *pagerduty) UpdateService(ctx context.Context, payload *types.CreateAndUpdateServicePayload) (*types.Service, error) {
	resp, err := pd.restClient.WithEndpoint(types.EndpointServices).
		WithToken(pd.token).
		WithBody(payload).
		PUT().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetServiceResp(resp)
	if err != nil {
		return nil, err
	}
	return getResp.Service, nil
}

func (pd *pagerduty) DeleteService(ctx context.Context, id string) error {
	_, err := pd.restClient.WithEndpoint(types.EndpointServices).
		WithToken(pd.token).
		AddPath(id).
		DELETE().
		Do(ctx)
	return err
}

func (pd *pagerduty) respToGetServiceResp(resp []byte) (getResp *types.GetServiceResp, err error) {
	err = json.Unmarshal(resp, &getResp)
	return getResp, err
}
