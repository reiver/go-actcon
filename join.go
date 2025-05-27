package actcon

import (
	"github.com/reiver/go-json"
)

type Join struct {
	Type json.Const[string] `json:"type" json.value:"Join"`

	Actor   string `json:"actor,omitempty"`   // who joined
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`  // what was joined
	Summary string `json:"summary,omitempty"` // for preview
}

func (receiver Join) String() string {
	return stringify(receiver)
}
