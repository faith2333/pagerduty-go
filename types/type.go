package types

const (
	TypeUser = "user"
)

type Type string

func (t Type) String() string {
	return string(t)
}
