package event

const (
	ActionTrigger     Action = "trigger"
	ActionAcknowledge Action = "acknowledge"
	ActionResolve     Action = "resolve"
)

type Action string

func (a Action) String() string {
	return string(a)
}
