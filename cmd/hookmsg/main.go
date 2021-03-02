package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IceWreck/HookMsg/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var routerMap map[string]http.Handler = make(map[string]http.Handler)

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

	// this is a clever trick (if i say so myself) to build only the router
	// which is needed for functionality mentioned in build tags (read Makefile)
	r.Mount("/hooks", routerMap["r"])

	log.Println("Running at Port ", config.Config.DeploymentPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.DeploymentPort), r)
	if err != nil {
		log.Println(err)
	}
}
