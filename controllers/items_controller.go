package controllers

import (
	"common_go/oauth"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"itemsModule/domain/items"
	"itemsModule/services"
	"itemsModule/utils/http_utils"
	"net/http"
	"strings"
	"utils/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	if oauth.GetCallerId(r) == 0 {
		respErr := rest_errors.NewUnauthorizedError("unauthorized")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, _ := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("Invalid json")
		http_utils.RespondError(w, respErr)
		return
	}
	itemRequest.Seller = oauth.GetCallerId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	fmt.Println(result)
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	result, getErr := services.ItemsService.Get(itemId)
	if getErr != nil {
		http_utils.RespondError(w, getErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, result)
}
