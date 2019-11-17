package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewSesseion try create new session
func NewSesseion(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	erro := !checkCredentials(*creds)

	if erro {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
