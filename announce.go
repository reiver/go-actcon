package actcon

import (
	"github.com/reiver/go-json"
)

type Announce struct {
	Type json.Const[string] `json:"type" json.value:"Announce"`

	Actor   string `json:"actor,omitempty"`   // who announced it
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`
	Summary string `json:"summary,omitempty"` // for preview
}

func (receiver Announce) String() string {
	return stringify(receiver)
}
