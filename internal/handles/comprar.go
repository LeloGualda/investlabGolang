package handles

import (
	"net/http"
)

// Comprar connection
func comprar(w http.ResponseWriter, r *http.Request) {
	returnStruct(w, r, queryGetCompraAcoes())
}

func getAcoes(w http.ResponseWriter, r *http.Request, id int) {
	returnStruct(w, r, queryGetAcoesUsuario(id))
}
func comprasLances(w http.ResponseWriter, r *http.Request, codigo string, id_user int) {
	returnStruct(w, r, queryGetLancesAcao(codigo, id_user))
}

func compraTemporalAcao(w http.ResponseWriter, r *http.Request, codigo string) {
	returnStruct(w, r, queryGetAcoesTemporal(codigo))
}
