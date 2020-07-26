package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Lgdev07/crud_api/models"
	"github.com/Lgdev07/crud_api/responses"
	"github.com/gorilla/mux"
)

func (a *App) createStore(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Venue successfully created"}

	store := &models.Store{}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &store)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	store.Prepare()

	storeCreated, err := store.Save(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["store"] = storeCreated
	responses.JSON(w, http.StatusCreated, resp)
	return

}

func (a *App) listStore(w http.ResponseWriter, r *http.Request) {
	store := &models.Store{}

	stores, err := store.GetStores(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, stores)
	return
}

func (a *App) updateStore(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{"status": "success"}

	params := mux.Vars(r)

	intID, _ := strconv.Atoi(params["id"])

	store, _ := models.GetStoreByID(intID, a.DB)

	if store.ID == 0 {
		resp["status"] = "failed"
		resp["message"] = "Store not found"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	updatedStore := models.Store{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if err := json.Unmarshal(body, &updatedStore); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	updatedStore.Prepare()

	_, err = updatedStore.UpdateStore(intID, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["message"] = "Store Updated"
	resp["store"] = updatedStore
	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) deleteStore(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"status": "Successed Request",
	}

	params := mux.Vars(r)

	intID, _ := strconv.Atoi(params["id"])

	store, _ := models.GetStoreByID(intID, a.DB)

	if store.ID == 0 {
		resp["status"] = "falied"
		resp["message"] = "Store not found"
		responses.JSON(w, http.StatusInternalServerError, resp)
		return
	}

	if err := models.DeleteStore(intID, a.DB); err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	resp["message"] = "Deleted Store " + params["id"]
	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) showStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	store, err := models.GetStoreByID(id, a.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, store)
	return
}
