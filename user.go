package pagerduty_go

import (
	"context"
	"encoding/json"
	"github.com/faith2333/pagerduty-go/types"
	"github.com/pkg/errors"
)

func (pd *pagerduty) GetUser(ctx context.Context, id string) (*types.User, error) {
	resp, err := pd.restClient.WithEndpoint(types.EndpointUsers).WithToken(pd.token).
		AddPath(id).
		GET().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getUserResp, err := pd.respToGetUserResp(resp)
	if err != nil {
		return nil, err
	}

	return getUserResp.User, nil
}

func (pd *pagerduty) CreateUser(ctx context.Context, userPayload *types.CreateAndUpdateUserPayload) (*types.User, error) {
	resp, err := pd.restClient.WithEndpoint(types.EndpointUsers).
		WithToken(pd.token).
		WithBody(userPayload).
		POST().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getUserResp, err := pd.respToGetUserResp(resp)
	if err != nil {
		return nil, err
	}

	return getUserResp.User, nil
}

// DeleteUser Returns 400 if the user has assigned incidents unless your pricing plan has the offboarding
// feature and the account is configured appropriately.
func (pd *pagerduty) DeleteUser(ctx context.Context, userID string) error {
	_, err := pd.restClient.WithEndpoint(types.EndpointUsers).
		AddPath(userID).
		WithToken(pd.token).
		DELETE().
		Do(ctx)
	return err
}

func (pd *pagerduty) UpdateUser(ctx context.Context, payload *types.CreateAndUpdateUserPayload) (*types.User, error) {
	resp, err := pd.restClient.WithEndpoint(types.EndpointUsers).
		WithToken(pd.token).
		WithBody(payload).
		PUT().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getUserResp, err := pd.respToGetUserResp(resp)
	if err != nil {
		return nil, err
	}

	return getUserResp.User, nil
}

func (pd *pagerduty) respToGetUserResp(resp []byte) (*types.GetUserResp, error) {
	getUserResp := &types.GetUserResp{}
	err := json.Unmarshal(resp, &getUserResp)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response to get user response failed")
	}

	return getUserResp, nil
}
