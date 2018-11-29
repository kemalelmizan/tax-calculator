package controller

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

// ResponseStruct ...
type ResponseStruct struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		root := ResponseStruct{
			Success: true,
			Data:    "Hello, World!",
		}

		// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(root)
	})
}

func ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&Healthy) == 1 {

			pong := ResponseStruct{
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
