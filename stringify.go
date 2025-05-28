package actsock

import (
	"github.com/reiver/go-json"
)

func stringify(value any) string {
	bytes, err := json.Marshal(value)
	if nil != err {
		return "﹝error﹞"
	}
	return string(bytes)
}
