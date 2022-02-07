package middleware

import (
	client "games-api/db"
	db "games-api/db/sqlc"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Fungsi Log yang berguna sebagai middleware
func AuthAdmin(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password Tidak boleh Kosong"))
			return
		}
		user, err := client.DB.GetUserByUsernameAndPassword(r.Context(), db.GetUserByUsernameAndPasswordParams{
			Username: uname,
			Password: pwd,
		})
		if err != nil {
			w.Write([]byte("Username atau Password tidak sesuai"))
			return
		}

		if user.Username == "admin" {
			next(w, r, ps)
		} else {
			w.Write([]byte("Anda bukan admin"))
			return
		}
	}
}

func AuthReviewer(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password Tidak boleh Kosong"))
			return
		}
		user, err := client.DB.GetUserByUsernameAndPassword(r.Context(), db.GetUserByUsernameAndPasswordParams{
			Username: uname,
			Password: pwd,
		})
		if err != nil {
			w.Write([]byte("Username atau Password tidak sesuai"))
			return
		}

		if user.Role == "reviewer" || user.Role == "admin" {
			next(w, r, ps)
		} else {
			w.Write([]byte("Anda bukan Reviewer"))
			return
		}
	}
}
