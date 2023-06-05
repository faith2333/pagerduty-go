package types

type Acknowledger struct {
}

type Acknowledgement struct {
	At           string              `json:"at"`
	Acknowledger BaseObjectReference `json:"acknowledger"`
}
