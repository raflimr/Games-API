package controller

import (
	"context"
	"encoding/json"
	"fmt"
	client "games-api/db"
	db "games-api/db/sqlc"
	"games-api/utils"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//POST
func PostRating(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var rating db.CreateRatingParams

		if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
			log.Println(err)
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.CreateRating(ctx, rating); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

// Update
func UpdateRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var rating db.UpdateRatingParams

		if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdateRating(ctx, rating); err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
