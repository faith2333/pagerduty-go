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

func TestPagerduty_ListUserContactMethods(t *testing.T) {
	contactMethods, err := testPagerDuty.ListUserContactMethods(context.Background(), "YOUR_UID")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	if len(contactMethods) == 0 {
		t.Log("resp is empty")
	}
}
