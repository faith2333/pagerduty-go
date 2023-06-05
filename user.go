package pagerduty_go

import (
	"context"
	"encoding/json"
	"fmt"
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

func (pd *pagerduty) ListUserContactMethods(ctx context.Context, uid string) ([]*types.ContactMethod, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointUsers).
		AddPath(pd.makeUserContactMethodPath(uid)).
		GET().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	listResp, err := pd.respToListUserContactMethodResp(resp)
	if err != nil {
		return nil, err
	}
	return listResp.ContactMethods, nil
}

func (pd *pagerduty) GetUserContactMethod(ctx context.Context, uid, contactMethodID string) (*types.ContactMethod, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointUsers).
		AddPath(pd.makeUserContactMethodPath(uid)).
		AddPath(contactMethodID).
		GET().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetUserContactMethodResp(resp)
	if err != nil {
		return nil, err
	}
	return getResp.ContactMethod, nil
}

func (pd *pagerduty) CreateUserContactMethod(ctx context.Context, uid string, payload *types.CreateAndUpdateContactMethodPayload) (*types.ContactMethod, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointUsers).
		AddPath(pd.makeUserContactMethodPath(uid)).
		WithBody(payload).
		POST().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetUserContactMethodResp(resp)
	if err != nil {
		return nil, err
	}
	return getResp.ContactMethod, nil
}

func (pd *pagerduty) UpdateUserContactMethod(ctx context.Context, uid, contactMethodID string, payload *types.CreateAndUpdateContactMethodPayload) (*types.ContactMethod, error) {
	resp, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointUsers).
		AddPath(pd.makeUserContactMethodPath(uid)).
		AddPath(contactMethodID).
		WithBody(payload).
		PUT().
		Do(ctx)
	if err != nil {
		return nil, err
	}

	getResp, err := pd.respToGetUserContactMethodResp(resp)
	if err != nil {
		return nil, err
	}
	return getResp.ContactMethod, nil
}

func (pd *pagerduty) DeleteUserContactMethod(ctx context.Context, uid, contactMethodID string) error {
	_, err := pd.restClient.WithToken(pd.token).
		WithEndpoint(types.EndpointUsers).
		AddPath(pd.makeUserContactMethodPath(uid)).
		AddPath(contactMethodID).
		DELETE().
		Do(ctx)
	return err
}

func (pd *pagerduty) respToGetUserResp(resp []byte) (*types.GetUserResp, error) {
	getUserResp := &types.GetUserResp{}
	err := json.Unmarshal(resp, &getUserResp)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response to get user response failed")
	}

	return getUserResp, nil
}

func (pd *pagerduty) makeUserContactMethodPath(uid string) string {
	return fmt.Sprintf("%s/contact_methods", uid)
}

func (pd *pagerduty) respToGetUserContactMethodResp(resp []byte) (getResp *types.GetContactMethodResp, err error) {
	err = json.Unmarshal(resp, &getResp)
	return
}

func (pd *pagerduty) respToListUserContactMethodResp(resp []byte) (listResp *types.ListContactMethodsResp, err error) {
	err = json.Unmarshal(resp, &listResp)
	return
}
