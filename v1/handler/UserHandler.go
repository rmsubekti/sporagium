package handler

import (
	"net/http"

	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/service"
)

func (v V1Handler) UserProfile(w http.ResponseWriter, r *http.Request) {
	var (
		user models.User
		err  error
	)
	logedInUser, err := v.getPrincipal(r)

	if err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if user, err = service.NewUserService().FirstById(logedInUser.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.JSON(w, http.StatusOK, user)
}
