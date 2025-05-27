package actcon

import (
	"github.com/reiver/go-json"
)

type Announce struct {
	Type json.Const[string] `json:"type" json.value:"Announce"`

	Object any `json:"object,omitempty"`
}

func (receiver Announce) String() string {
	return stringify(receiver)
}
