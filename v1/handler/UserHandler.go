package handler

import (
	"net/http"

	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/model"
)

func (v V1Handler) UserProfile(w http.ResponseWriter, r *http.Request) {
	var (
		User model.User
	)
	logedInUser, err := v.getPrincipal(r)

	if err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := User.Get(logedInUser.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.JSON(w, http.StatusOK, User)
}
