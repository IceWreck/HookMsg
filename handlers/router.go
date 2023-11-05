package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *Application) routes() http.Handler {
	r := chi.NewRouter()

	// custom error handlers
	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! Welcome to Anchit's HookMsg Service."))
	})

	r.Get("/healthcheck", app.healthCheck)

	r.Route("/hooks", func(r chi.Router) {
		r.Get("/script/{endpoint}", app.scriptHook)
		r.Post("/script/{endpoint}", app.scriptHook)

		if app.config.TelegramEnabled {
			r.Post("/telegram", app.telegramHook)
		}

		if app.config.MatrixEnabled {
			r.Post("/matrix/{channel}", app.matrixHook)
		}

	})

	return r
}
