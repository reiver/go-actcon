package actsock

import (
	"github.com/reiver/go-json"
)

type Leave struct {
	Type json.Const[string] `json:"type" json.value:"Leave"`

	Actor   string `json:"actor,omitempty"`   // who joined
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`  // what was joined
	Summary string `json:"summary,omitempty"` // for preview
}

func (receiver Leave) String() string {
	return stringify(receiver)
}
