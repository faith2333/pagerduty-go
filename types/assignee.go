package types

import "time"

type Assignee struct {
}

type Assignment struct {
	At       time.Time           `json:"at"`
	Assignee BaseObjectReference `json:"assignee"`
}
