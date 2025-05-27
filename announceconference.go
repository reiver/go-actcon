package actcon

import (
	"github.com/reiver/go-json"
)

type AnnounceConference struct {
	Type json.Const[string] `json:"type" json.value:"Announce"`
	Object Conference       `json:"object,omitempty"`
}
