package dto

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, code int, data any) {
	response := Response{
		Code:   code,
		Status: http.StatusText(code),
		Data:   data,
	}

	dataJson, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dataJson)
}
