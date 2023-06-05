package pagerduty_go

import (
	"context"
	"testing"
)

func TestPagerduty_GetUser(t *testing.T) {
	user, err := testPagerDuty.GetUser(context.Background(), "123")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if user == nil {
		t.Log("response in nil")
		t.FailNow()
	}
}
