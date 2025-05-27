package actcon

import (
	"github.com/reiver/go-json"
)

type Invite struct {
	Type json.Const[string] `json:"type" json.value:"Invite"`

	Actor   string `json:"actor,omitempty"`   // who did the inviting (i.e., the inviter)
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`    // for preview
	Object  any    `json:"object,omitempty"`  // what the inviter invited the invitee to
	Summary string `json:"summary,omitempty"` // for preview
	Target  string `json:"target,omitempty"`  // who was invite (i.e., the invitee)
}
