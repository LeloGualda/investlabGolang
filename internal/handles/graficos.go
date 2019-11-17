package handles

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// GraficH Gerencia todos os graficos
func GraficH(w http.ResponseWriter, r *http.Request) {
	// auth, _ := checkUser(r)
	// if !auth {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	// sid := strings.TrimPrefix(r.URL.Path, "/grafic/")
	// println(sid)

	// w.Header().Set("Content-Type", "application/json")
	// if sid == "saldo" {
	// 	saldo(w, r)
	// }
}
