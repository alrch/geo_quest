package api

import (
	"geo_quest/internal/api/handler"
	"geo_quest/internal/service"
	"net/http"
)

func NewRouter(generator service.QuestGenerator) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/quest", &handler.QuestHandler{Generator: generator})
	return mux
}
