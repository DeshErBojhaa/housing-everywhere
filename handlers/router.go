package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DeshErBojhaa/housing-everywhere/web"
)

// NewAPI returns the api configured to handle all our requests.
func NewAPI(log *log.Logger, sd chan os.Signal) http.Handler {
	app := web.NewApp(log, sd)
	loc := Location{log: log}

	//! Assumption!! if drones send the sector id along with the location, we can have a central
	//! service that can cover all the requests.
	app.Handle("GET", "/v1/atlas/dns/location/:sector_id", loc.Atlas)
	app.Handle("GET", "/v1/mama/dns/location/:sector_id", loc.Mama)
	return app
}
