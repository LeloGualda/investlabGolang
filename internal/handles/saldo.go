package handles

import (
	"net/http"

	"../querys"
)

func saldo(w http.ResponseWriter, r *http.Request, id int) {
	returnStruct(w, r, querys.QueryGetSaldo(id))
}
