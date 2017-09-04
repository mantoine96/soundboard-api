package soundboard

import (
	"net/http"
	"soundboard-api/logger"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route : structure to define a HTTP Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes for our API
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"AddSound",
		"POST",
		"/",
		controller.AddSound,
	},
	Route{
		"UpdateSound",
		"PUT",
		"/",
		controller.UpdateSound,
	},
	Route{
		"DeleteSound",
		"DELETE",
		"/",
		controller.DeleteSound,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
