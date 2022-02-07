package controller

import (
	"context"
	"encoding/json"
	"fmt"
	client "games-api/db"
	db "games-api/db/sqlc"
	"games-api/utils"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//GET
func GetImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		idInt, _ := strconv.Atoi(id)

		image, err := client.DB.GetGambar(ctx, int32(idInt))

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, image, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

//POST
func PostImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var gambar db.CreateGambarParams

		if err := json.NewDecoder(r.Body).Decode(&gambar); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		_, errorImageURL := url.ParseRequestURI(gambar.Url)
		var errorMessage struct {
			Error string `json:"error"`
		}
		if errorImageURL != nil {
			errorMessage.Error = "Image URL Tidak Valid"
			utils.ResponseJSON(w, errorMessage, http.StatusInternalServerError)
			return
		}

		if err := client.DB.CreateGambar(ctx, gambar); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}
