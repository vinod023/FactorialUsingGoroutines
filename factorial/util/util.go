package util

import (
	"encoding/json"
	"net/http"
)

//ResponseData model for displaying the status
type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

//RespondJSON return the http response in json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	res := ResponseData{
		Status: status,
		Data:   payload,
	}
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

//RespondError return the http method error
func RespondError(w http.ResponseWriter, status int, message string) {
	res := ResponseData{
		Status: status,
		Data:   message,
	}
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
