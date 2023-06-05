package pagerduty_go

import (
	"context"
	"github.com/faith2333/pagerduty-go/event"
	"github.com/faith2333/pagerduty-go/types"
)

type Interface interface {
	WithRESTClient(restClient IRESTClient) Interface
	IUser
	IService
	IIncident
	IEvent
}

type IUser interface {
	GetUser(ctx context.Context, id string) (*types.User, error)
	// CreateUser Create a new user.
	// Users are members of a PagerDuty account that have the ability to interact with Incidents and other data on the account.
	//      Scoped OAuth requires: users.write
	//  reference: https://developer.pagerduty.com/api-reference/4cb4fd0f5444a-create-a-user
	CreateUser(ctx context.Context, payload *types.CreateAndUpdateUserPayload) (*types.User, error)
	// UpdateUser Update an existing user.
	// Users are members of a PagerDuty account that have the ability to interact with Incidents and other data on the account.
	//      Scoped OAuth requires: users.write
	//  reference: https://developer.pagerduty.com/api-reference/ce6799fc6191a-update-a-user
	UpdateUser(ctx context.Context, uid string, payload *types.CreateAndUpdateUserPayload) (*types.User, error)
	// DeleteUser Remove an existing user.
	// Returns 400 if the user has assigned incidents unless your pricing plan has the offboarding feature and the account is configured appropriately.
	// Note that the incidents reassignment process is asynchronous and has no guarantee to complete before the api call return.
	// Users are members of a PagerDuty account that have the ability to interact with Incidents and other data on the account.
	//      Scoped OAuth requires: users.write
	//  reference: https://developer.pagerduty.com/api-reference/f99c2c2bba70b-delete-a-user
	DeleteUser(ctx context.Context, id string) error

	// ListUserContactMethods
	//   reference: https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods
	ListUserContactMethods(ctx context.Context, uid string) ([]*types.ContactMethod, error)
	// GetUserContactMethod
	//   reference: https://developer.pagerduty.com/api-reference/e210330b7a2fb-get-a-user-s-contact-method
	GetUserContactMethod(ctx context.Context, uid, contactMethodID string) (*types.ContactMethod, error)
	// CreateUserContactMethod
	//   reference: https://developer.pagerduty.com/api-reference/38d473dce8f6b-create-a-user-contact-method
	CreateUserContactMethod(ctx context.Context, uid string, payload *types.CreateAndUpdateContactMethodPayload) (*types.ContactMethod, error)
	// UpdateUserContactMethod
	//   reference: https://developer.pagerduty.com/api-reference/d5732eccb6f3c-update-a-user-s-contact-method
	UpdateUserContactMethod(ctx context.Context, uid, contactMethodID string, payload *types.CreateAndUpdateContactMethodPayload) (*types.ContactMethod, error)
	// DeleteUserContactMethod
	//   reference: https://developer.pagerduty.com/api-reference/d5732eccb6f3c-update-a-user-s-contact-method
	DeleteUserContactMethod(ctx context.Context, uid, contactMethodID string) error
}

type IService interface {
	// GetService
	//   reference: https://developer.pagerduty.com/api-reference/165ad96a22ffd-get-a-service
	GetService(ctx context.Context, id string) (*types.Service, error)
	// CreateService
	//   reference: https://developer.pagerduty.com/api-reference/7062f2631b397-create-a-service
	CreateService(ctx context.Context, payload *types.CreateAndUpdateServicePayload) (*types.Service, error)
	// UpdateService
	//   reference: https://developer.pagerduty.com/api-reference/fbc6e9f4ef8eb-update-a-service
	UpdateService(ctx context.Context, payload *types.CreateAndUpdateServicePayload) (*types.Service, error)
	// DeleteService
	//   reference: https://developer.pagerduty.com/api-reference/fbc6e9f4ef8eb-update-a-service
	DeleteService(ctx context.Context, id string) error
}

type IIncident interface {
	// ListIncidents
	//  reference: https://developer.pagerduty.com/api-reference/9d0b4b12e36f9-list-incidents
	ListIncidents(ctx context.Context, params *types.ListIncidentsReq) (*types.ListIncidentsResp, error)
	// GetIncident
	//  reference: https://developer.pagerduty.com/api-reference/005299ed43553-get-an-incident
	GetIncident(ctx context.Context, iid string) (*types.Incident, error)
	// CreateIncident
	//  reference: https://developer.pagerduty.com/api-reference/a7d81b0e9200f-create-an-incident
	CreateIncident(ctx context.Context, payload *types.CreateIncidentPayload) (*types.Incident, error)
}

type IEvent interface {
	SendEvent(ctx context.Context, payload *event.SendEventReq) (*event.SendEventResp, error)
}
