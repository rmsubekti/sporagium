package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/middleware"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/service"

	session "github.com/go-session/session/v3"
)

func (v V1Handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		user models.User
	)
	userSrv := service.NewUserService()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := userSrv.Register(&user); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = ""
	dto.JSON(w, http.StatusOK, user)
}

func (v V1Handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		login dto.Login
		user  models.User
		err   error
	)
	store, _ := session.Start(r.Context(), w, r)
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	userSrv := service.NewUserService()
	if user, err = userSrv.Login(login); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	principal := middleware.Principal{
		ID:         user.ID.String(),
		Name:       user.Name,
		ExpireDays: 7,
	}

	if err := principal.CreateToken(); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	store.Set("J5E", principal.Token)
	store.Set("U5E", principal.ID)
	store.Save()

	if _, ok := store.Get("ReturnUri"); ok {
		dto.JSON(w, http.StatusTemporaryRedirect, "/auth")
		return
	}

	dto.JSON(w, http.StatusOK, principal)

}
