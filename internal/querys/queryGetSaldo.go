package querys

import (
	"fmt"

	"../structs"
)

func QueryGetSaldo(id_user int) structs.Transacao {
	u := &structs.Transacao{}
	err := db.QueryRow("select sum(valor) from mydb.transacao where id_user = ?", id_user).Scan(&u.Valor)
	if err != nil {
		fmt.Println(err.Error())
	}
	return *u
}
