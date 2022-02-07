package controller

import (
	"context"
	"encoding/json"
	"fmt"
	client "games-api/db"
	db "games-api/db/sqlc"
	"games-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//GET
func GetGame(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		games, err := client.DB.ListGames(ctx)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, games, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

//POST
func PostGames(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var games db.CreateGamesParams

		if err := json.NewDecoder(r.Body).Decode(&games); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.CreateGames(ctx, games); err != nil {
			if strings.Contains(err.Error(), "foreign key constraint fails") {
				kesalahan := map[string]string{
					"error": "Data Developer/Publisher tidak ada",
				}

				utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
				return
			}
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
func UpdateGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var games db.UpdateGamesParams

		if err := json.NewDecoder(r.Body).Decode(&games); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdateGames(ctx, games); err != nil {
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

// Delete
func DeleteGames(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		idInt, _ := strconv.Atoi(id)
		if err := client.DB.DeleteGames(ctx, int32(idInt)); err != nil {
			if strings.Contains(err.Error(), "foreign key constraint fails") {
				kesalahan := map[string]string{
					"error": "Data tidak dapat dihapus, masih ada data yang terikat",
				}

				utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
				return
			}
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
