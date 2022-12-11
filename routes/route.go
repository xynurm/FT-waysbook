package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	BookRoutes(r)
	CartRoutes(r)
	TransactionRoutes(r)
}