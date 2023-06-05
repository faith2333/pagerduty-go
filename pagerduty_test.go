package pagerduty_go

import (
	"os"
	"testing"
)

var testPagerDuty Interface

func TestMain(m *testing.M) {
	testPagerDuty = NewPagerDuty("y_NbAkKc66ryYTWUXYEu")
	os.Exit(m.Run())
}
