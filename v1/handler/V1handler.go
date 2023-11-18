package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/rmsubekti/sporagium/middleware"
)

type V1Handler struct {
}

func (v V1Handler) getPrincipal(r *http.Request) (principal middleware.Principal, err error) {
	var ok bool
	ctx := context.Get(r, "principal")
	if principal, ok = ctx.(middleware.Principal); !ok {
		err = fmt.Errorf("you cannot access this resource")
	}
	return
}
