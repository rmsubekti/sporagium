package handler

import (
	"encoding/json"
	"net/http"

	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/middleware"
	"github.com/rmsubekti/sporagium/model"

	session "github.com/go-session/session/v3"
)

func (v V1Handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		register dto.Register
		account  model.Account
	)

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := account.Set(register).Create(); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	register.Password = ""
	dto.JSON(w, http.StatusOK, register)
}

func (v V1Handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		login   dto.Login
		account model.Account
	)
	store, _ := session.Start(r.Context(), w, r)
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := account.Login(login.Email, login.Password); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	principal := middleware.Principal{
		ID:         account.ID.String(),
		Name:       account.User.Name,
		ExpireDays: 7,
	}
	if err := principal.CreateToken(); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	store.Set("J5E", principal.Token)
	store.Set("U5E", account.ID)
	store.Save()
	dto.JSON(w, http.StatusOK, principal)

}
