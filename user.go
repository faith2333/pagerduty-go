package pagerduty_go

import (
	"context"
	"fmt"
	"github.com/faith2333/pagerduty-go/types"
)

func (pd *pagerduty) GetUser(ctx context.Context, id string) (*types.User, error) {
	resp, err := pd.restClient.WithEndpoint(types.EndpointUsers).WithToken(pd.token).GET(id).Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(resp))

	return nil, nil
}
