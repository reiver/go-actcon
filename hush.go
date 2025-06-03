package actsock

import (
	"github.com/reiver/go-json"
)

type Hush struct {
	Type json.Const[string] `json:"type" json.value:"Hush"`

	Value  bool   `json:"@value"`
}

func (receiver Hush) String() string {
	return stringify(receiver)
}
