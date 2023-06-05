package types

const (
	NotificationRuleTypeDefault NotificationRuleType = "assignment_notification_rule_reference"
)

type NotificationRuleType string

func (nr NotificationRuleType) String() string {
	return string(nr)
}

type NotificationRule struct {
	Type NotificationRuleType `json:"type"`
	Base
}
