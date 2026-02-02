package handler

import (
	"encoding/json"
	"geo_quest/internal/service"
	"net/http"
	"strconv"
)

type QuestHandler struct {
	Generator service.QuestGenerator
}

func (h *QuestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("lat")
	lngStr := r.URL.Query().Get("lng")
	style := r.URL.Query().Get("style")
	difficulty := r.URL.Query().Get("difficulty")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "invalid latitude", http.StatusBadRequest)
		return
	}
	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		http.Error(w, "invalid longitude", http.StatusBadRequest)
		return
	}

	quest, err := h.Generator.GenerateQuest(lat, lng, style, difficulty)
	if err != nil {
		http.Error(w, "failed to generate quest", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quest)
}
