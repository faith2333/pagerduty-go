package pagerduty_go

import (
	"context"
	"github.com/faith2333/pagerduty-go/types"
)

type Interface interface {
	GetUser(ctx context.Context, id string) (*types.User, error)
	CreateUser(ctx context.Context, payload *types.CreateAndUpdateUserPayload) (*types.User, error)
	UpdateUser(ctx context.Context, payload *types.CreateAndUpdateUserPayload) (*types.User, error)
	DeleteUser(ctx context.Context, id string) error
}
