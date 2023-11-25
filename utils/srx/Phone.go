package srx

import (
	"errors"
	"regexp"
)

type Phone string

var rxPhone = regexp.MustCompile(`^(?:[[:digit:]]+)$`)

func (p Phone) Validate() (err error) {
	if !rxPhone.MatchString(string(p)) {
		return errors.New("not valid phone number")
	}
	return
}

func (p Phone) Ok() bool {
	return rxPhone.MatchString(string(p))
}
