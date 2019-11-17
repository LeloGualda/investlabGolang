package handles

import (
	"net/http"

	"../querys"
)

// Comprar connection
func comprar(w http.ResponseWriter, r *http.Request) {
	returnStruct(w, r, querys.QueryGetCompraAcoes())
}

func getAcoes(w http.ResponseWriter, r *http.Request, id int) {
	returnStruct(w, r, querys.QueryGetAcoesUsuario(id))
}
func comprasLances(w http.ResponseWriter, r *http.Request, codigo string, id_user int) {
	returnStruct(w, r, querys.QueryGetLancesAcao(codigo, id_user))
}

func compraTemporalAcao(w http.ResponseWriter, r *http.Request, codigo string) {
	returnStruct(w, r, querys.QueryGetAcoesTemporal(codigo))
}
