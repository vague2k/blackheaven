package handlers

import "github.com/vague2k/blackheaven/db"

type Handler struct {
	DB *db.DB
}

func NewHandler() *Handler {
	db, err := db.Init("")
	if err != nil {
		panic(err)
	}

	return &Handler{
		DB: db,
	}
}
