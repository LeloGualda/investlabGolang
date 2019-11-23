package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// UserH send user info or get per user
func UserH(w http.ResponseWriter, r *http.Request) {

	auth, u := CheckUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sid := strings.TrimPrefix(r.URL.Path, "/user/")
	println(sid)

	w.Header().Set("Content-Type", "application/json")
	if sid == "info" {
		json, _ := json.Marshal(u)
		fmt.Fprint(w, string(json))
	}
}
