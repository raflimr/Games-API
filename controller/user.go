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
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//POST
func PostUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var user db.CreateUserParams

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := client.DB.CreateUser(ctx, user); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
	}
}

//GET
func GetUsersByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		ambilId := ps.ByName("id")
		ambilIdInt, _ := strconv.Atoi(ambilId)
		defer cancel()
		var user db.User
		user, err := client.DB.GetUserByID(ctx, int32(ambilIdInt))

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, user, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

//GET
func GetUsersByUsernameAndPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		var users db.CreateUserParams

		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			log.Println(err)
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		user, err := client.DB.GetUserByUsernameAndPassword(ctx, db.GetUserByUsernameAndPasswordParams{
			Username: users.Username,
			Password: users.Password,
		})

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, user, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

//GET
func GetAdminUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()
		var users db.CreateUserParams

		if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
			log.Println(err)
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		user, err := client.DB.GetAdminUser(ctx)

		if err != nil {
			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}
			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		utils.ResponseJSON(w, user, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}
