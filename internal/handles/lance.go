package handles

import (
	"net/http"

	"../querys"
)

func lance(w http.ResponseWriter, r *http.Request, qtd, lance, id_user int, venda bool, codigo string) {

	acoes := querys.QueryGetAcoesCarteiraUser(id_user, qtd, codigo, venda)

	for _, acao := range acoes {

		if venda {
			acao.Lance = lance
		} else {
			acao.Lance = 0
		}

		acao.Venda = venda
		querys.QueryUpdateAcoesCarteira(acao)
	}
}
