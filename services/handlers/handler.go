package handlers

import (
	"github.com/vague2k/blackheaven/services/database"
)

type Handler struct {
	DB *database.Queries
}

func NewHandler() *Handler {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	return &Handler{DB: db}
}
