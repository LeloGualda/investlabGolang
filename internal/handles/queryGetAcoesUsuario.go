package handles

import (
	"fmt"
	"strconv"
)

type AcoesDoUsuario struct {
	Codigo string `json:"codigo", db:"codigo"`
	Lance  int    `json:"lance", db:"lance"`
	Qtd    int    `json:"qtd", db:"qtd"`
	Venda  bool   `json:"venda", db:"venda"`
}

func queryGetAcoesUsuario(id int) []AcoesDoUsuario {

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
					mydb.valores.data = (select max(data) from mydb.valores group by codigo)) as a
						group by codigo,venda,lance;
	`

	var acoes = []AcoesDoUsuario{}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		acao := &AcoesDoUsuario{}

		rows.Scan(&acao.Codigo, &acao.Venda, &acao.Lance, &acao.Qtd)

		acoes = append(acoes, *acao)
	}
	return acoes
}
