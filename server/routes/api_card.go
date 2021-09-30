package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Joraffe/jobine_tcg/server/services/db"
)


/*---------------------------------
          Response Data
----------------------------------*/
// CardIDResponseData is the data we send
// back after a successfully creating a new card
type CardIDResponseData struct {
	CardID  int64  `json:"cardId"`
}


// CardResponseData is the data we send
// for fetching an individual card
type CardResponseData struct {
	Card  *db.Card  `json:"card"`
}


// CardsResponseData is the data we send back
// after a successfully getting all card data
type CardsResponseData struct {
	Cards  []*db.Card  `json:"cards"`
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
	// GET Request Handlers
	case http.MethodGet:
    switch head {
		case "all":
		  r.handleGetAll(res, req)
		case "":
			queryParams := req.URL.Query()
			if param, ok := queryParams["type"]; ok {
				r.handleGetByType(res, req, strings.Join(param, ""))
				return
			}
			if param, ok := queryParams["name"]; ok {
				r.handleGetByName(res, req, strings.Join(param, ""))
				return
			}
		default:
			http.Error(res, fmt.Sprintf("Unsupported GET path %s", head), http.StatusBadRequest)
			return
		}
	// POST Request Handlers
	case http.MethodPost:
		// TODO: Add "admin role" type permissions system
		switch head {
		case "create":
			r.handleCreate(res, req)
		case "update":
			r.handleUpdate(res, req)
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

	cardID, err := r.Services.Database.CreateCard(cardCreate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error creating new card in database: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success: true,
		Error:   nil,
		Data:    CardIDResponseData{
			CardID:  cardID,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}


func (r *CardRouter) handleUpdate(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	cardUpdate := new(db.CardUpdate)

	err := decoder.Decode(cardUpdate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid JSON request: %s", err.Error()), http.StatusBadRequest)
		return
	}

	cardID, err := r.Services.Database.UpdateCard(cardUpdate)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error updating card in database: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success:   true,
		Error:     nil,
		Data:      CardIDResponseData{
			CardID:  cardID,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}


func (r *CardRouter) handleGetAll(res http.ResponseWriter, req *http.Request) {
	cards, err := r.Services.Database.GetAllCards()
	if err != nil {
		http.Error(res, fmt.Sprintf("Error getting all cards from DB: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success:  true,
		Error:    nil,
		Data:     CardsResponseData{
			Cards:  cards,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}


func (r *CardRouter) handleGetByType(res http.ResponseWriter, req *http.Request, typeParam string) {
	cards, err := r.Services.Database.GetCardsByCardType(typeParam)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error getting all cards for given type from DB: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success:  true,
		Error:    nil,
		Data:     CardsResponseData{
			Cards:  cards,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}


func (r *CardRouter) handleGetByName(res http.ResponseWriter, req *http.Request, nameParam string) {
	cards, err := r.Services.Database.GetCardsByCardType(nameParam)
	if err != nil {
		http.Error(res, fmt.Sprintf("Error getting card with name %s from DB: %s", nameParam, err.Error()), http.StatusInternalServerError)
		return
	}

	response := &Response{
		Success:  true,
		Error:    nil,
		Data:     CardsResponseData{
			Cards:   cards,
		},
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(response)
}
