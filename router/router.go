package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"score/common"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(
	healthController common.IHealthController,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	DefineHealthRoutes(router, healthController)
	return router
}

func DefineHealthRoutes(router *mux.Router, healthController common.IHealthController) {
	router.Methods("GET").Path("/health").HandlerFunc(healthController.GetHealth)
}
