package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Lgdev07/crud_api/models"
	"github.com/Lgdev07/crud_api/responses"
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
	stores, err := models.GetStores(&a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, stores)
	return
}

// func (a *App) updateStore(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	body, err := ioutil.ReadAll(r.Body)

// 	defer r.Body.Close()

// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	type ParamsProps struct {
// 		Name string `json:"name"`
// 	}

// 	var paramsProps ParamsProps

// 	err = json.Unmarshal(body, &paramsProps)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	params := mux.Vars(r)

// 	for _, store := range stores {
// 		if store.ID == params["id"] {
// 			store.Name = paramsProps.Name
// 			output, err := json.Marshal(store)

// 			if err != nil {
// 				http.Error(w, err.Error(), 500)
// 				return
// 			}

// 			w.Write(output)
// 			return
// 		}
// 	}
// }

// func (a *App) deleteStore(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	for index, value := range stores {
// 		if value.ID == params["id"] {
// 			stores[index] = stores[len(stores)-1]
// 			stores = stores[:len(stores)-1]
// 			json.NewEncoder(w).Encode(stores)
// 			return
// 		}
// 	}
// }

// func (a *App) listStore(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(stores)
// }

// func (a *App) showStore(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)
// 	for _, item := range stores {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// }
