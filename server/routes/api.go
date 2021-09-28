package routes

import (
	"net/http"
)


/*---------------------------------
             Responses
----------------------------------*/
// Response is the data 
type Response struct {
  Success  bool         `json:"success"`
  Error    error        `json:"error"`
  Data     interface{}  `json:"data"`
}


/*---------------------------------
             Router
----------------------------------*/
// APIRouter is responsible for serving "/api"
// or delegating to the appropriate sub api-router
type APIRouter struct {
	Services    *Services
	CardRouter  *CardRouter
	SetRouter   *SetRouter 
}


func (r *APIRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)

	// TODO: Add auth for API routes
	// if head == "auth" {
  //   r.AuthRouter.ServeHTTP(res, req)
  //   return
  // }

	switch head {
	case "card":
	  r.CardRouter.ServeHTTP(res, req)
	case "set":
		r.SetRouter.ServeHTTP(res, req)
	default:
		http.Error(res, "404 Not Found", http.StatusNotFound)
	}
}


// NewAPIRouter makes a new api router and sets up its children
// routers with access to our router services
func NewAPIRouter(routerServices *Services) *APIRouter {
	router := new(APIRouter)

	router.Services = routerServices
	router.CardRouter = NewCardRouter(routerServices)
	router.SetRouter = NewSetRouter(routerServices)

	return router
}
