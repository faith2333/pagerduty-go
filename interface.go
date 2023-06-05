package pagerduty_go

import (
	"context"
	"github.com/faith2333/pagerduty-go/types"
)

type Interface interface {
	GetUser(ctx context.Context, id string) (*types.User, error)
	// CreateUser Create a new user.
	// Users are members of a PagerDuty account that have the ability to interact with Incidents and other data on the account.
	//      Scoped OAuth requires: users.write
	//  https://developer.pagerduty.com/api-reference/4cb4fd0f5444a-create-a-user
	CreateUser(ctx context.Context, payload *types.CreateAndUpdateUserPayload) (*types.User, error)
	// UpdateUser Update an existing user.
	// Users are members of a PagerDuty account that have the ability to interact with Incidents and other data on the account.
	//      Scoped OAuth requires: users.write
	//  https://developer.pagerduty.com/api-reference/ce6799fc6191a-update-a-user
	UpdateUser(ctx context.Context, payload *types.CreateAndUpdateUserPayload) (*types.User, error)
	// DeleteUser Remove an existing user.
	// Returns 400 if the user has assigned incidents unless your pricing plan has the offboarding feature and the account is configured appropriately.
	// Note that the incidents reassignment process is asynchronous and has no guarantee to complete before the api call return.
	// Users are members of a PagerDuty account that have the ability to interact with Incidents and other data on the account.
	//      Scoped OAuth requires: users.write
	// https://developer.pagerduty.com/api-reference/f99c2c2bba70b-delete-a-user
	DeleteUser(ctx context.Context, id string) error
}
