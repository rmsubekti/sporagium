package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/middleware"
	"github.com/rmsubekti/sporagium/model"
)

func (v V1Handler) CreateSpora(w http.ResponseWriter, r *http.Request) {
	var spora model.Spora
	user, err := v.getPrincipal(r)
	if err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&spora); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := spora.Create(user.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.JSON(w, http.StatusOK, spora)
}

func (v V1Handler) ViewListSpora(w http.ResponseWriter, r *http.Request) {
	var sporas model.Sporas
	user, err := v.getPrincipal(r)
	if err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := sporas.GetAll(user.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	dto.JSON(w, http.StatusOK, sporas)
}

func (v V1Handler) CreateClientSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var spora model.Spora
	var client model.Client
	var err error
	var user middleware.Principal
	if user, err = v.getPrincipal(r); err != nil {
		dto.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}

	if err = spora.Get(id); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err)
		return
	}

	if spora.UserID != uuid.MustParse(user.ID) {
		dto.JSON(w, http.StatusUnauthorized, "youre not the owner of this spora")
	}

	client.Domain = spora.CallbackURL
	client.SporaID = spora.ID
	client.Secret = uuid.NewString()
	if err = client.Create(); err != nil {
		dto.JSON(w, http.StatusInternalServerError, spora)
		return
	}
	dto.JSON(w, http.StatusOK, "client secret created")
}
