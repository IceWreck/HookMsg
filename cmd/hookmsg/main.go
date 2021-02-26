package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IceWreck/HookMsg/config"
	"github.com/IceWreck/HookMsg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	// Initialize Router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! Welcome to Anchit's HookMsg Service."))
	})

	r.Route("/hooks", func(r chi.Router) {
		r.Post("/script/{endpoint}", handlers.ScriptHook)
		// r.Post("/telegram", handlers.TelegramHook)
		r.Post("/matrix/{channel}", handlers.MatrixHook)
	})

	log.Println("Running at Port ", config.Config.DeploymentPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.DeploymentPort), r)
	if err != nil {
		log.Println(err)
	}
}
