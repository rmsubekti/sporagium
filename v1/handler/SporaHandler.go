package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rmsubekti/sporagium/dto"
	"github.com/rmsubekti/sporagium/helper"
	"github.com/rmsubekti/sporagium/middleware"
	"github.com/rmsubekti/sporagium/models"
	"github.com/rmsubekti/sporagium/service"
)

func (v V1Handler) CreateSpora(w http.ResponseWriter, r *http.Request) {
	var spora models.Spora
	user, err := v.getPrincipal(r)
	if err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&spora); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	sporaSrv := service.NewSporaService()
	if err := sporaSrv.Create(&spora, user.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.JSON(w, http.StatusOK, spora)
}

func (v V1Handler) ViewListSpora(w http.ResponseWriter, r *http.Request) {
	var paginator helper.Paginator
	user, err := v.getPrincipal(r)
	if err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&paginator); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	sporaSrv := service.NewSporaService()
	if err := sporaSrv.Paginate(&paginator, user.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.JSON(w, http.StatusOK, paginator)
}

func (v V1Handler) GenerateSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var spora models.Spora
	var err error
	var user middleware.Principal
	if user, err = v.getPrincipal(r); err != nil {
		dto.JSON(w, http.StatusUnauthorized, err.Error())
		return
	}
	if spora, err = service.NewSporaService().FirstByID(id); err != nil {
		dto.JSON(w, http.StatusNotFound, err)
		return
	}
	if spora.UserID != uuid.MustParse(user.ID) {
		dto.JSON(w, http.StatusUnauthorized, "youre not the owner of this spora")
	}
	if err = service.NewSecretService().Generate(spora.ID); err != nil {
		dto.JSON(w, http.StatusInternalServerError, err)
	}
	dto.JSON(w, http.StatusOK, "client secret created")
}
