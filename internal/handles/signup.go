package handles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../querys"
	"../structs"
)

// Signup create new user
func Signup(w http.ResponseWriter, r *http.Request) {

	creds := &structs.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = querys.CreateUser(*creds)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	transVenda := new(structs.Transacao)
	transVenda.ID = querys.QueryGetUseID(creds.Username)
	transVenda.Tipo = 1
	transVenda.Valor = 200000
	transVenda.Descricao = "deposito inicial"
	querys.InsertTransacao(*transVenda)

}
