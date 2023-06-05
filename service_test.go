package pagerduty_go

import (
	"context"
	"fmt"
	"github.com/faith2333/pagerduty-go/types"
	"testing"
)

func TestPagerduty_GetService(t *testing.T) {
	svc, err := testPagerDuty.GetService(context.TODO(), "id")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if svc.ID != "id" || svc.Type != types.TypeService {
		t.Log(fmt.Sprintf("response failed except: {id:\"id\", type:\"service\"}, but got: %#v", svc))
		t.FailNow()
	}
}
