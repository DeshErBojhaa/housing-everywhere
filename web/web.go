package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"syscall"

	"github.com/dimfeld/httptreemux"
	"go.opencensus.io/plugin/ochttp"
)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*httptreemux.TreeMux
	log *log.Logger

	// OpenCensus http is used for future extension on tracing.
	och      *ochttp.Handler
	shutdown chan os.Signal
}

// NewApp is constructor for our application backend.
func NewApp(log *log.Logger, sd chan os.Signal) *App {
	app := App{
		TreeMux:  httptreemux.New(),
		log:      log,
		shutdown: sd,
	}

	app.och = &ochttp.Handler{
		Handler: app.TreeMux,
	}
	return &app
}

// Handler is the signature of out custom handlers.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error

// Handle handels incoming http requests
func (a *App) Handle(verb, path string, handler Handler) {
	// This function to execute on each request
	h := func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		if err := handler(r.Context(), w, r, params); err != nil {
			log.Printf("XXXX critical shutdown error %v", err)
			a.SignalShutdown()
		}
	}
	a.TreeMux.Handle(verb, path, h)
}

// SignalShutdown ...
func (a *App) SignalShutdown() {
	a.log.Printf("error from handler caused integrity issue. causing shutdown")
	a.shutdown <- syscall.SIGSTOP

}

// ServeHTTP implements the http.Handler interface. It overrides the ServeHTTP
// of the embedded TreeMux by using the ochttp.Handler instead. That Handler
// wraps the TreeMux handler so the routes are served.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.och.ServeHTTP(w, r)
}
