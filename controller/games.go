package controller

import (
	"encoding/json"
	"gle/models"
	"net/http"
)

type Games struct{}

func (g Games) Index(w http.ResponseWriter, r *http.Request) {
	games := models.GetAllGames()

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(games)
}
