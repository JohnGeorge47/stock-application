package handlers

import (
	"encoding/json"
	"github.com/JohnGeorge47/stock-application/internal/core/user"
	"net/http"
)

type Validate struct {
	Method  string `json:"method"`
	Success bool   `json:"success"`
	Valid   bool   `json:"valid"`
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var response Validate
		response.Method = http.MethodGet
		session_token := r.Header.Get("session-token")
		if session_token == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		qparams := r.URL.Query()["email_id"]
		val, err := user.Validate(session_token, qparams[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Success = true
		response.Valid = *val
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(responseJson)
	}
}
