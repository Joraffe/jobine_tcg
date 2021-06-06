package routes

import (
	"net/http"
)


// AppRouter is the router responsible for serving "/"
// or delegating to the appropriate sub router
type AppRouter struct {
  StaticRouter  *StaticRouter
}

func (r *AppRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)

	switch head {
	case "static":
		r.StaticRouter.ServeHTTP(res, req)
	default:
		http.ServeFile(res, req, "dist/static/index.html")
	}
}


// NewRouter makes a new app router and sets up its children
// routers with access to the router services
func NewRouter() *AppRouter {
  router := new(AppRouter)
  router.StaticRouter = NewStaticRouter()

  return router
}
