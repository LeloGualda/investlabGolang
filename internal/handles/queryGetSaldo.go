package handles

import "fmt"

type Transacao struct {
	ID        int     `json:"id", db:"id_transacao"`
	Data      string  `json:"data", db:"data"`
	Valor     float64 `json:"valor", db:"valor"`
	Tipo      int     `json:"tipo", db:"tipo"`
	Descricao string  `json:"descricao", db:"descricao"`
}

func queryGetSaldo(id_user int) Transacao {
	u := &Transacao{}
	err := db.QueryRow("select sum(valor) from mydb.transacao where id_user = ?", id_user).Scan(&u.Valor)
	if err != nil {
		fmt.Println(err.Error())
	}
	return *u
}
