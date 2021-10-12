package handlers

import (
	"net/http"
	"time"

	"github.com/IceWreck/HookMsg/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.Application) http.Handler {
	r := chi.NewRouter()

	// custom error handlers
	r.NotFound(notFoundResponse(app))
	r.MethodNotAllowed(methodNotAllowedResponse(app))

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

	r.Get("/healthcheck", healthCheck(app))

	r.Route("/hooks", func(r chi.Router) {
		r.Post("/script/{endpoint}", scriptHook(app))

		if app.Config.TelegramEnabled {
			r.Post("/telegram", telegramHook(app))
		}

		if app.Config.MatrixEnabled {
			r.Post("/matrix/{channel}", matrixHook(app))
		}

	})

	return r
}
