package event

const (
	SeverityCritical Severity = "critical"
	SeverityWarning  Severity = "warning"
	SeverityError    Severity = "error"
	SeverityInfo     Severity = "info"
)

type Severity string

func (s Severity) String() string {
	return string(s)
}
