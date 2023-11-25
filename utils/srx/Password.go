package srx

import (
	"errors"
	"regexp"
)

type Password string

var upper = regexp.MustCompile(`[[:upper:]]+`)
var lower = regexp.MustCompile(`[[:lower:]]+`)
var punct = regexp.MustCompile(`[[:punct:]]+`)
var digit = regexp.MustCompile(`[[:digit:]]+`)

func (p Password) Validate() (err error) {
	if !upper.MatchString(string(p)) {
		err = errors.New("password should contain one or more uppercase character")
	}
	if !lower.MatchString(string(p)) {
		err = errors.New("password should contain one or more lowercase character")
	}
	if !punct.MatchString(string(p)) {
		err = errors.New("password should contain one or more special character")
	}
	if !digit.MatchString(string(p)) {
		err = errors.New("password should contain one or more digit character")
	}
	return
}

func (p Password) Ok() bool {
	if !upper.MatchString(string(p)) {
		return false
	}
	if !lower.MatchString(string(p)) {
		return false
	}
	if !punct.MatchString(string(p)) {
		return false
	}
	return digit.MatchString(string(p))
}
