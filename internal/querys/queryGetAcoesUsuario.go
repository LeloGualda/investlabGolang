package querys

import (
	"fmt"
	"strconv"

	"../structs"
)

func QueryGetAcoesUsuario(id int) []structs.AcoesDoUsuario {

	query := `
	select codigo,
	venda,
	lance,
	count(codigo)  as qtd
	from	(select
			mydb.carteira.id_carteira,
			mydb.carteira.id_user,
			mydb.carteira.codigo,
			mydb.carteira.venda,
			mydb.carteira.lance
			from mydb.carteira,mydb.valores
				where
					mydb.carteira.id_user = ` + strconv.Itoa(id) + `
						and
					mydb.valores.data = (select max(data) from valores
											where
												valores.codigo = carteira.codigo
											group by codigo)) as a
						group by codigo,venda,lance;
	`

	var acoes = []structs.AcoesDoUsuario{}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		acao := &structs.AcoesDoUsuario{}

		rows.Scan(&acao.Codigo, &acao.Venda, &acao.Lance, &acao.Qtd)

		acoes = append(acoes, *acao)
	}
	return acoes
}
