package main

import (
	"context"
	"database/sql"
	"fmt"
	"games-api/config"
	"games-api/controller"
	client "games-api/db"
	sqlClient "games-api/db/sqlc"
	"games-api/middleware"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func CekAdmin() {
	admin, err := client.DB.GetAdminUser(context.TODO())
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	if admin.ID == 0 {
		client.DB.CreateUser(context.TODO(), sqlClient.CreateUserParams{
			Username: "admin",
			Password: "admin",
			Role:     "admin",
		})
	}
}

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	client.DB = sqlClient.New(db)
	CekAdmin()
	fmt.Println("Success")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := httprouter.New()
	router.POST("/user/register", controller.PostUsers)
	router.POST("/user/login", controller.GetUsersByUsernameAndPassword)
	router.GET("/user/cek-user/:id", controller.GetUsersByID)
	router.GET("/publisher", controller.GetPublisher)
	router.POST("/publisher", middleware.AuthAdmin(controller.PostPublisher))
	router.PUT("/publisher", middleware.AuthAdmin(controller.UpdatePublisher))
	router.DELETE("/publisher", middleware.AuthAdmin(controller.DeletePublisher))
	router.GET("/developer", controller.GetDeveloper)
	router.POST("/developer", middleware.AuthAdmin(controller.PostDeveloper))
	router.PUT("/developer", middleware.AuthAdmin(controller.UpdateDeveloper))
	router.DELETE("/developer", middleware.AuthAdmin(controller.DeleteDeveloper))
	router.GET("/games", controller.GetGame)
	router.POST("/games", middleware.AuthAdmin(controller.PostGames))
	router.PUT("/games", middleware.AuthAdmin(controller.UpdateGames))
	router.DELETE("/games", middleware.AuthAdmin(controller.DeleteGames))
	router.GET("/images", controller.GetImage)
	router.POST("/images", middleware.AuthAdmin(controller.PostImage))
	router.GET("/kategori-game", controller.GetKategoriGames)
	router.POST("/kategori-game", middleware.AuthAdmin(controller.PostKategoriGame))
	router.GET("/kategori", controller.GetKategori)
	router.POST("/kategori", middleware.AuthAdmin(controller.PostKategori))
	router.PUT("/kategori", middleware.AuthAdmin(controller.UpdateKategori))
	router.DELETE("/kategori", middleware.AuthAdmin(controller.DeleteKategori))
	router.POST("/rating", middleware.AuthReviewer(controller.PostRating))
	router.PUT("/rating", middleware.AuthReviewer(controller.UpdateRating))
	router.GET("/review", controller.GetReview)
	router.POST("/review", middleware.AuthReviewer(controller.PostReview))
	router.PUT("/review", middleware.AuthReviewer(controller.UpdateReview))
	router.DELETE("/review", middleware.AuthReviewer(controller.DeleteReview))
	router.GET("/reaksi-review", controller.GetReaksiReview)
	router.POST("/reaksi-review", middleware.AuthReviewer(controller.PostReaksiReview))
	router.PUT("/reaksi-review", middleware.AuthReviewer(controller.UpdateReaksiReview))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
