package pagerduty_go

import (
	"os"
	"testing"
)

var testPagerDuty Interface

func TestMain(m *testing.M) {
	testPagerDuty = NewPagerDuty("y_NbAkKc66ryYTWUXYEu").
		WithRESTClient(NewDefaultRestClient().WithHost("https://stoplight.io/mocks/pagerduty-upgrade/api-schema/2748099"))
	os.Exit(m.Run())
}
