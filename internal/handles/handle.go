package handles

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Api(w http.ResponseWriter, r *http.Request) {
	auth, user := checkUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sid := strings.TrimPrefix(r.URL.Path, "/api/")

	if strings.HasPrefix(sid, "admin") {
		sid = strings.TrimPrefix(sid, "admin/")
		if strings.HasPrefix(sid, "search/") {
			keyword, ok := r.URL.Query()["keyword"]
			if !ok || len(keyword[0]) < 1 {
				fmt.Println("Url Param 'keyword' is missing")
				return
			}

			adminSearchAPI(w, r, keyword[0])
		}

		if strings.HasPrefix(sid, "addCarteira/") {

		}

	}

	if strings.HasPrefix(sid, "comprar") {
		if sid == "comprar" {
			comprar(w, r)
			return
		}

		sid = strings.TrimPrefix(sid, "comprar/")

		if strings.HasPrefix(sid, "temporal/") {
			sid = strings.TrimPrefix(sid, "temporal/")
			compraTemporalAcao(w, r, sid)
			return
		}

		fmt.Println(sid)
		comprasLances(w, r, sid, user.ID)
	}

	if strings.HasPrefix(sid, "acoes") {
		if sid == "acoes" {
			getAcoes(w, r, user.ID)
			return
		}
	}

	if strings.HasPrefix(sid, "saldo") {
		if sid == "saldo" {
			saldo(w, r, user.ID)
			return
		}
	}

	if strings.HasPrefix(sid, "lance") {

		lancesqtd, ok := r.URL.Query()["qtd"]

		if !ok || len(lancesqtd[0]) < 1 {
			fmt.Println("Url Param 'vender' is missing")
			return
		}

		lanceVender, ok := r.URL.Query()["vender"]
		if !ok || len(lanceVender[0]) < 1 {
			fmt.Println("Url Param 'vender' is missing")
			return
		}

		lanceValor, ok := r.URL.Query()["lance"]

		if !ok || len(lanceValor[0]) < 1 {
			fmt.Println("Url Param 'valor do lance' is missing")
			return
		}

		lanceCodigo, ok := r.URL.Query()["codigo"]

		if !ok || len(lanceCodigo[0]) < 1 {
			fmt.Println("Url Param 'valor do lance' is missing")
			return
		}

		lanceQtd, _ := strconv.Atoi(lancesqtd[0])
		lanceValorConv, _ := strconv.Atoi(lanceValor[0])
		lanceVenda, _ := strconv.ParseBool(lanceVender[0])

		lance(w, r, lanceQtd, lanceValorConv, user.ID, lanceVenda, lanceCodigo[0])
	}

	if strings.HasPrefix(sid, "vender") {

		sid = strings.TrimPrefix(sid, "vender/")
		valorR, ok1 := r.URL.Query()["valor"]
		codigoR, ok := r.URL.Query()["codigo"]
		qtdR, ok2 := r.URL.Query()["qtd"]

		if !ok || len(codigoR[0]) < 1 {
			fmt.Println("Url Param 'codigo' is missing")
			return
		}
		if !ok1 || len(valorR[0]) < 1 {
			fmt.Println("Url Param 'valor' is missing")
			return
		}

		if !ok2 || len(qtdR[0]) < 1 {
			fmt.Println("Url Param 'qtd' is missing")
			return
		}

		codigo := codigoR[0]
		valor := valorR[0]
		qtd := qtdR[0]

		fmt.Println(codigo)
		fmt.Println(valor)
		fmt.Println(qtd)
		v, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			return
		}
		i, err := strconv.Atoi(qtd)
		if err != nil {
			return
		}
		venderAcao(w, r, v, codigo, user.Username, i)
	}
}
