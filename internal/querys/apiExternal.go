package querys

import (
	"fmt"

	"../structs"
)

func InsertAcao(acao structs.ApiAcao) {
	query := "INSERT INTO acoes (codigo,nome,regiao,tipo,moeda) VALUES (?,?,?,?,?);"
	_, err := db.Query(query, acao.Codigo, acao.Name, acao.Region, acao.Type, acao.Currency)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InsertCarteira(carteira structs.Carteira) {
	// fmt.Println(carteira.Codigo,)
	query := "INSERT INTO carteira (id_user,codigo,venda,lance) VALUES (?,?,?,?);"
	_, err := db.Query(query, carteira.ID, carteira.Codigo, carteira.Venda, carteira.Lance)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func InsertValorAcao(valor structs.BaseValores) {
	query := "INSERT INTO valores (codigo,data,valor) VALUES (?,?,?);"

	_, err := db.Query(query, valor.Codigo, valor.Date, valor.Valor)

	if err != nil {
		fmt.Println(err.Error())
	}
}
