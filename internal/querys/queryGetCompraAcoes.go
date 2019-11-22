package querys

import "fmt"

type BaseValores struct {
	Codigo string  `json:"codigo",db:"codigo"`
	Date   string  `json:"data",db:"data"`
	Valor  float64 `json:"valor",db:"valor"`
}

func QueryGetCompraAcoes() []BaseValores {

	query := `
	select
	data,
	valor,
	codigo
	from mydb.valores as a
		where
	data = (select max(data) from valores as b
    where b.codigo = a.codigo
    group by codigo);

`

	var valores = []BaseValores{}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		valor := &BaseValores{}

		rows.Scan(&valor.Date, &valor.Valor, &valor.Codigo)

		valores = append(valores, *valor)
	}
	return valores
}

func QueryGetAcoesTemporal(codigo string) []BaseValores {

	query := `
	SELECT valores.data,valores.valor,valores.codigo  FROM mydb.valores
		where codigo = ?
			order by valores.data
				desc limit 10;
`

	var valores = []BaseValores{}

	rows, err := db.Query(query, codigo)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		valor := &BaseValores{}

		rows.Scan(&valor.Date, &valor.Valor, &valor.Codigo)

		valores = append(valores, *valor)
	}
	return valores
}
