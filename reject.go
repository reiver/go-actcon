package actcon

import (
	"github.com/reiver/go-json"
)

type Reject struct {
	Type json.Const[string] `json:"type" json.value:"Reject"`

	Actor   string `json:"actor,omitempty"`   // who did the rejecting
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`  // the thing that is being rejected
	Summary string `json:"summary,omitempty"` // for preview
}
