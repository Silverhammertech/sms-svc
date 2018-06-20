package rest

import (
	"net/http"
	"net/http/pprof"
	"log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	swagger "github.com/Silverhammertech/swagger-lib"
	oauth2 "github.com/Silverhammertech/oauth-lib"
)

const default_port = "8080"
const base_path = "/api/v1/"

func attachProfiler(router *mux.Router) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

func getRouter() (router *mux.Router) {

	// Register a handler for each route pattern
	router = mux.NewRouter()

	// Add a trivial handler for INFO
	attachProfiler(router)

	// attach swagger documentation api
	err := swagger.AttachSwaggerUI(router, base_path)
	if err != nil {
		log.Panic(err.Error())
	}

	//  standard endpoints
	api := router.PathPrefix(base_path).Subrouter()

	//  these should not require authentication to get results
	api.Path("/ping").Methods("GET").HandlerFunc(HandlePing)

	// Custom REST handlers
	oauth2.SetOauthState(false) // TODO: set oauth state
	for _, route := range routes {
		api.Path(route.Pattern).Methods(route.Method).Handler(oauth2.AuthHandler(route.HandlerFunc))
	}

	return
}

func StartServer() {

	// TODO: get dynamic port from server
	port := default_port

	// get router object
	router := getRouter()

	// Start listening on the configured port.
	// ListenAndServe is not expected to return, so we wrap it in a log.Fatal
	// also include CORS handling
	err:= http.ListenAndServe(":"+port, handlers.CORS()(router))
	if(err != nil){
		log.Panic(err.Error())
	}
}
