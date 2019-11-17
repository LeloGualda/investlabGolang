package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Signup create new user
func Signup(w http.ResponseWriter, r *http.Request) {

	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createUser(*creds)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
