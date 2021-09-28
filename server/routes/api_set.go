package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Joraffe/jobine_tcg/server/services/db"
)


/*---------------------------------
          Response Data
----------------------------------*/
// SetCreateResponseData is the data we send
// back after a successfully creating a new set
type SetCreateResponseData struct {
	Set  *db.Set  `json:"set"`
}


/*---------------------------------
             Router
----------------------------------*/
// SetRouter is responsible for serving "/api/set"
// Basically, connecting to our Postgres DB for all
// of the CRUD operations for our "Set" models
type SetRouter struct {
	Services *Services
}


func (r *SetRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)

	switch req.Method {
	// POST Request Handlers
	case http.MethodPost:
		// TODO: Add "admin role" type permissions system
		switch head {
		case "create":
			r.handleCreate(res, req)
		default:
			http.Error(res, fmt.Sprintf("Unsupported POST path %s", head), http.StatusBadRequest)
			return
		}
	// Unsupported Method Response
	default:
		http.Error(res, fmt.Sprintf("Unsupported Method type %s", req.Method), http.StatusBadRequest)
	}
}


// NewSetRouter makes a new api/set router and hooks up its services
func NewSetRouter(routerServices *Services) *SetRouter {
	router := new(SetRouter)

	router.Services = routerServices

	return router
}


/*---------------------------------
             Handlers
----------------------------------*/
func (r *SetRouter) handleCreate(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	setCreate := new(db.SetCreate)

	err := decoder.Decode(setCreate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid JSON request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	set, err := r.Services.Database.CreateSet(setCreate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error creating new set in database: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success:  true,
		Error:    nil,
		Data:     SetCreateResponseData{
			Set:  set,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
