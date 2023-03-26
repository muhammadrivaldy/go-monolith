package util

import "gopkg.in/guregu/null.v4"

func NewString(s string) (ns null.String) {
	if s != "" {
		ns = null.StringFrom(s)
	}

	return
}
