package main

import (
	//...

	"WebMsg/actions"
	"WebMsg/hooks"
	"WebMsg/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Initialize Telegram
	go actions.InitializeTG()

	// Initialize Router
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! Welcome to Anchit's WebMsg Microservice."))
	})

	r.Route("/hooks", func(r chi.Router) {
		r.Post("/telegram", hooks.TelegramHook)
		r.Post("/script/{endpoint}", hooks.ScriptHook)
	})

	log.Println("Running it at port ", utils.Config.DeploymentPort)
	http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.DeploymentPort), r)
}
