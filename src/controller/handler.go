package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

// Response ...
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}

// PostBill ...
func PostBill() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		type ProductInputWrapper struct {
			Data []ProductInput `json:"data"`
		}
		var productInputWrapper ProductInputWrapper

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &productInputWrapper)
		if err != nil {
			panic(err)
		}

		root := Response{
			Success: true,
			Data:    PostProduct(productInputWrapper.Data),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(root)
	})
}

// Index ...
func Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	})
}

// Ping ...
func Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if atomic.LoadInt32(&Healthy) == 1 {

			pong := Response{
				Success: true,
				Data:    "pong",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(pong)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}
