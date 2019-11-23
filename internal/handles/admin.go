package handles

import (
	"net/http"

	"../alphavantage"
	"../querys"
	"../structs"
)

func adminSearchAPI(w http.ResponseWriter, r *http.Request, keyword string) {
	returnStruct(w, r, alphavantage.GetAPi(keyword, false))
}

func adminAddAcao(w http.ResponseWriter, r *http.Request, codigo, nome, tipo string) {
	el := new(structs.ApiAcao)
	el.Codigo = codigo
	el.Name = nome
	el.Type = tipo
	querys.InsertAcao(*el)
	alphavantage.InserirSeriesAPI(codigo)
}

func updateSeries(codigo string) {
	alphavantage.InserirSeriesAPI(codigo)
}

func ativarAcao(codigo string) {
	alphavantage.AddAcaoBot(codigo)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	returnStruct(w, r, querys.QueryGetUsers())
}
