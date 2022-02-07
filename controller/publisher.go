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
	"net/url"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//GET
func GetPublisher(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		publisher, err := client.DB.ListPublisher(ctx)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, publisher, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

//POST
func PostPublisher(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var publisher db.CreatePublisherParams

		if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
			log.Println(err)
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		_, errorImageURL := url.ParseRequestURI(publisher.Logopub)
		var errorMessage struct {
			Error string `json:"error"`
		}
		if errorImageURL != nil {
			errorMessage.Error = "Image URL Tidak Valid"
			utils.ResponseJSON(w, errorMessage, http.StatusInternalServerError)
			return
		}

		_, errorImageURLWebsite := url.ParseRequestURI(publisher.Website)
		if errorImageURLWebsite != nil {
			errorMessage.Error = "Website URL Tidak Valid"
			utils.ResponseJSON(w, errorMessage, http.StatusInternalServerError)
			return
		}

		if err := client.DB.CreatePublisher(ctx, publisher); err != nil {
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
func UpdatePublisher(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "PUT" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var publisher db.UpdatePublisherParams

		if err := json.NewDecoder(r.Body).Decode(&publisher); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.UpdatePublisher(ctx, publisher); err != nil {
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
func DeletePublisher(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		idInt, _ := strconv.Atoi(id)
		if err := client.DB.DeletePublisher(ctx, int32(idInt)); err != nil {
			if strings.Contains(err.Error(), "foreign key constraint fails") {
				kesalahan := map[string]string{
					"error": "Publisher mempunyai relasi dengan table game sehingga data publisher tidak dapat dihapus ",
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
