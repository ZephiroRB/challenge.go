package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ZephiroRB/challenge.go/api/auth"
	"github.com/ZephiroRB/challenge.go/api/models"
	"github.com/ZephiroRB/challenge.go/api/responses"
	"github.com/ZephiroRB/challenge.go/api/utils/formaterror"
)

func (server *Server) CreateSkill(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	skill := models.Skill{}
	err = json.Unmarshal(body, &skill)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	skill.Prepare()
	err = skill.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	log.Print("uid '%s'", uid)

	skillCreated, err := skill.SaveSkill(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, skillCreated.ID))
	responses.JSON(w, http.StatusCreated, skillCreated)
}
