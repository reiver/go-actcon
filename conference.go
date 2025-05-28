package actsock

import (
	"github.com/reiver/go-json"
)

// Conference is an 'Object' of 'type' 'Conference'.
type Conference struct {
	Type json.Const[string] `json:"type" json.value:"Conference"`

	Actor     string `json:"actor,omitempty"`
	Closed    string `json:"closed,omitempty"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Origin  []string `json:"origin,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	Summary   string `json:"summary,omitempty"`
	To      []string `json:"to,omitempty"`
}

func (receiver Conference) String() string {
	return stringify(receiver)
}
