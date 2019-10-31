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

	app.Handle("GET", "/v1/mumma/location/:id")
	return app

}
