package handles

import "net/http"

func lance(w http.ResponseWriter, r *http.Request, qtd, lance, id_user int, venda bool, codigo string) {

	acoes := queryGetAcoesCarteiraUser(id_user, qtd, codigo, venda)

	for _, acao := range acoes {

		if venda {
			acao.Lance = lance
		} else {
			acao.Lance = 0
		}

		acao.Venda = venda
		queryUpdateAcoesCarteira(acao)
	}
}
