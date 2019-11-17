package querys

import (
	"fmt"

	"../structs"
)

func QueryUpdateAcoesCarteira(acao structs.AcoesDoUsuario) bool {

	query := `
	UPDATE carteira SET
 		venda = ?,
 		lance = ?
		WHERE id_carteira = ?;
`
	//TODO: USANDO QUANTIDADE COMO ID ALTERAR
	_, err := db.Query(query, acao.Venda, acao.Lance, acao.Qtd)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func QueryGetAcoesCarteiraUser(id_user, limit int, codigo string, venda bool) []structs.AcoesDoUsuario {

	query := `
	select codigo,venda,lance,id_carteira
		from carteira
		where
			carteira.id_user = ? and
			carteira.codigo = ? and
			carteira.venda = ?
		limit ?
`
	rows, err := db.Query(query, id_user, codigo, !venda, limit)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var acoes = []structs.AcoesDoUsuario{}

	for rows.Next() {
		acao := &structs.AcoesDoUsuario{}
		//TODO: USANDO QUANTIDADE COMO ID ALTERAR
		rows.Scan(&acao.Codigo, &acao.Venda, &acao.Lance, &acao.Qtd)

		acoes = append(acoes, *acao)
	}
	return acoes
}
