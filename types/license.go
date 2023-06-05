package types

const (
	LicenseTypeDefault = "license_reference"
)

type LicenseType string

func (lt LicenseType) String() string {
	return string(lt)
}

type License struct {
	Type LicenseType `json:"type"`
	Base
}
