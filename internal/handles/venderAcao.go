package handles

import (
	"fmt"
	"net/http"
)

func venderAcao(w http.ResponseWriter, r *http.Request, valor float64, codigo, user string, qtd int) {
	err := insertCompra(user, codigo, valor, qtd)

	if err != nil {
		fmt.Println(err.Error())
	}
}
