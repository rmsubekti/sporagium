package srx

import (
	"errors"
	"regexp"
)

type Email string

var rxMail = regexp.MustCompile(`^(?:[[:alnum:]]+[[:alnum:]\-\.]+[[:alnum:]])+@(?:[[:alnum:]]+[[:alnum:]\-\.]+[[:alnum:]])+\.(?:[[:alpha:]]{2,6})$`)

func (e Email) Validate() (err error) {
	if !rxMail.MatchString(string(e)) {
		return errors.New("not valid email")
	}
	return
}

func (e Email) Ok() bool {
	return rxMail.MatchString(string(e))
}
