package actcon

import (
	"github.com/reiver/go-json"
)

type Create struct {
	Type json.Const[string] `json:"type" json.value:"Create"`

	Actor   string `json:"actor,omitempty"`   // who did the creating (i.e., the creater)
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`  // what was created
	Summary string `json:"summary,omitempty"` // for preview
}
