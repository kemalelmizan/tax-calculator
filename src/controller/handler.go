package controller

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/kemalelmizan/tax-calculator/src/model"
)

// Response ...
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}

// PostBill ...
func PostBill(db *sql.DB, log *log.Logger) http.Handler {
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

		// TODO: pass DB and logger from main
		pm := model.NewProductModel(db, log)
		pc := NewProductController(pm)

		billOutput, err := pc.PostProduct(productInputWrapper.Data)
		if err != nil {
			panic(err)
		}

		root := Response{
			Success: true,
			Data:    billOutput,
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
