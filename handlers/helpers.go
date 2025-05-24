package handlers

import (
	"encoding/json"
	"net/http"
)

type RespErr struct {
	Status   int    `json:"status"`
	ErrorMsg string `json:"content"`
}

func respErr(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(&RespErr{
		Status:   status,
		ErrorMsg: errMsg,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
}
