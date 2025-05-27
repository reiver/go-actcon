package actcon

import (
	"github.com/reiver/go-json"
)

// Conference is an 'Object' of 'type' 'Conference'.
type Conference struct {
	Type json.Const[string] `json:"type" json.value:"Conference"`

	Actor     string `json:"actor,omitempty"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Origin  []string `json:"origin,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	Summary   string `json:"summary,omitempty"`
	To      []string `json:"to,omitempty"`
}
