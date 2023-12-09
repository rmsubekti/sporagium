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
		reg dto.Register
	)
	accountSrv := service.NewAccountService()
	if err := json.NewDecoder(r.Body).Decode(&reg); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	account := reg.GetAccount()
	if err := accountSrv.Register(&account); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	account.Password = ""
	dto.JSON(w, http.StatusOK, account.User)
}

func (v V1Handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		login   dto.Login
		account models.Account
		err     error
	)
	store, _ := session.Start(r.Context(), w, r)
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	accountSrv := service.NewAccountService()
	if account, err = accountSrv.Login(login); err != nil {
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
	store.Set("U5E", principal.ID)
	store.Save()

	if _, ok := store.Get("ReturnUri"); ok {
		dto.JSON(w, http.StatusTemporaryRedirect, "/auth")
		return
	}

	dto.JSON(w, http.StatusOK, principal)

}
