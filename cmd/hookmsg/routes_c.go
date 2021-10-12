package main

import (
	"github.com/IceWreck/HookMsg/handlers"
	"github.com/go-chi/chi"
)

func init() {
	r := chi.NewRouter()
	r.Post("/script/{endpoint}", handlers.ScriptHook)
	r.Post("/telegram", handlers.TelegramHook)
	r.Post("/matrix/{channel}", handlers.MatrixHook)
	routerMap["r"] = r
}
