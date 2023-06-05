package pagerduty_go

import (
	"os"
	"testing"
)

var testPagerDuty Interface

func TestMain(m *testing.M) {
	testPagerDuty = NewPagerDuty("u+nhy-LtvijP-S9nozzQ")
	os.Exit(m.Run())
}
