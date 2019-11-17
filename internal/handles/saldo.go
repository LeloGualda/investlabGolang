package handles

import (
	"net/http"
)

func saldo(w http.ResponseWriter, r *http.Request, id int) {
	returnStruct(w, r, queryGetSaldo(id))
}
