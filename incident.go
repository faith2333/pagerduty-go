package pagerduty_go

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/faith2333/pagerduty-go/types"
)

func (pd *pagerduty) ListIncidents(ctx context.Context, params *types.ListIncidentsReq) (*types.ListIncidentsResp, error) {
	client := pd.restClient.WithToken(pd.token).WithEndpoint(types.EndpointIncidents)

	reqMap, err := params.AsMap()
	if err != nil {
		return nil, err
	}

	for k, v := range reqMap {
		client.AddURLParam(k, fmt.Sprintf("%v", v))
	}

	resp, err := client.GET().Do(ctx)
	if err != nil {
		return nil, err
	}
	return pd.respToListIncidentsResp(resp)
}

func (pd *pagerduty) GetIncident(ctx context.Context, iid string) (*types.Incident, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointIncidents).
		AddPath(iid).
		GET().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetIncidentResp(resp)
	if err != nil {
		return nil, err
	}

	return getResp.Incident, nil
}

func (pd *pagerduty) CreateIncident(ctx context.Context, payload *types.CreateIncidentPayload) (*types.Incident, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointIncidents).
		WithBody(payload).
		POST().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetIncidentResp(resp)
	if err != nil {
		return nil, err
	}

	return getResp.Incident, nil
}

func (pd *pagerduty) respToListIncidentsResp(resp []byte) (listResp *types.ListIncidentsResp, err error) {
	err = json.Unmarshal(resp, &listResp)
	return
}

func (pd *pagerduty) respToGetIncidentResp(resp []byte) (getResp *types.GetIncidentPayload, err error) {
	err = json.Unmarshal(resp, &getResp)
	return
}
