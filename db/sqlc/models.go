// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Developer struct {
	ID             int32  `json:"id"`
	Logodev        string `json:"logodev"`
	KantorPusat    string `json:"kantor_pusat"`
	Pendiri        string `json:"pendiri"`
	TahunPendirian string `json:"tahun_pendirian"`
}

type Gambar struct {
	ID     int32  `json:"id"`
	Url    string `json:"url"`
	GameID int32  `json:"game_id"`
}

type Game struct {
	ID           int32  `json:"id"`
	Judul        string `json:"judul"`
	Deskripsi    string `json:"deskripsi"`
	Penerbit     string `json:"penerbit"`
	Platform     string `json:"platform"`
	Website      string `json:"website"`
	StatusGame   string `json:"status_game"`
	TanggalRilis string `json:"tanggal_rilis"`
	DeveloperID  int32  `json:"developer_id"`
	PublisherID  int32  `json:"publisher_id"`
}

type Kategori struct {
	ID        int32  `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

type KategoriGame struct {
	KategoriID int32 `json:"kategori_id"`
	GameID     int32 `json:"game_id"`
}

type Publisher struct {
	ID        int32  `json:"id"`
	Logopub   string `json:"logopub"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Website   string `json:"website"`
}

type Rating struct {
	ID     int32 `json:"id"`
	Rating int32 `json:"rating"`
	GameID int32 `json:"game_id"`
	UserID int32 `json:"user_id"`
}

type ReaksiReview struct {
	ID         int32  `json:"id"`
	TipeReaksi string `json:"tipe_reaksi"`
	ReviewID   int32  `json:"review_id"`
	UserID     int32  `json:"user_id"`
}

type Review struct {
	ID            int32     `json:"id"`
	KontenReview  string    `json:"konten_review"`
	TanggalReview time.Time `json:"tanggal_review"`
	Tipe          string    `json:"tipe"`
	GameID        int32     `json:"game_id"`
	UserID        int32     `json:"user_id"`
}

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
