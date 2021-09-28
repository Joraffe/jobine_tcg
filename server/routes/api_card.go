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
// CardCreateResponseData is the data we send
// back after a successfully creating a new card
type CardCreateResponseData struct {
	Card  *db.Card  `json:"card"`
}


/*---------------------------------
             Router
----------------------------------*/
// CardRouter is responsible for serving "/api/card"
// Basically, connecting to our Postgres DB for all
// of the CRUD operations for our "Card" models
type CardRouter struct {
	Services  *Services
}


func (r *CardRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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

// NewCardRouter makes a new api/card router and hooks up its services
func NewCardRouter(routerServices *Services) *CardRouter {
	router := new(CardRouter)

	router.Services = routerServices

	return router
}


/*---------------------------------
             Handlers
----------------------------------*/
func (r *CardRouter) handleCreate(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	cardCreate := new(db.CardCreate)

	err := decoder.Decode(cardCreate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid JSON request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	card, err := r.Services.Database.CreateCard(cardCreate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error creating new card in database: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success: true,
		Error:   nil,
		Data:    CardCreateResponseData{
			Card:  card,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
