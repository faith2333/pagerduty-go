package pagerduty_go

import (
	"os"
	"testing"
)

var testPagerDuty Interface

func TestMain(m *testing.M) {
	testPagerDuty = NewPagerDuty("YOUR_TOKEN")
	os.Exit(m.Run())
}
