package main

import (
	//...

	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IceWreck/HookMsg/hooks"
	"github.com/IceWreck/HookMsg/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Initialize Telegram
	// go actions.InitializeTG()

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
		//r.Post("/telegram", hooks.TelegramHook)
		r.Post("/script/{endpoint}", hooks.ScriptHook)
	})

	log.Println("Running it at port ", utils.Config.DeploymentPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.DeploymentPort), r)
	if err != nil {
		log.Println(err)
	}
}
