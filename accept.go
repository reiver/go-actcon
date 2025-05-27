package actcon

import (
	"github.com/reiver/go-json"
)

type Accept struct {
	Type json.Const[string] `json:"type" json.value:"Accept"`

	Actor   string `json:"actor,omitempty"`   // who did the accepting
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`  // the thing that is being accepted
	Summary string `json:"summary,omitempty"` // for preview
}

func (receiver Accept) String() string {
	return stringify(receiver)
}
