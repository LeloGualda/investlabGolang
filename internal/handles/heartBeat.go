package handles

import "net/http"

// Test connection
func Test(w http.ResponseWriter, r *http.Request) {
	auth, _ := CheckUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	println("Test connection ok")
}
