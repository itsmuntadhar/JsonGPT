package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	lib "github.com/itsmuntadhar/JsonGPT/pkg"
)

func Serve(port string, apiKey string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}
		defer r.Body.Close()
		var u lib.Request
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		u.APIKey = apiKey
		resp, err := lib.GetGPTResponse(u)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(resp))
	})
	fmt.Println("Server started on port " + port)
	http.ListenAndServe(":"+port, mux)
}
