package handles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

// Test connection
func Test(w http.ResponseWriter, r *http.Request) {
	auth, _ := checkUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	println("Test connection ok")
}

func Api(w http.ResponseWriter, r *http.Request) {
	auth, user := checkUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sid := strings.TrimPrefix(r.URL.Path, "/api/")

	if strings.HasPrefix(sid, "comprar") {
		if sid == "comprar" {
			comprar(w, r)
			return
		}

		sid = strings.TrimPrefix(sid, "comprar/")
		fmt.Println(sid)
		comprasLances(w, r, sid)
	}

	if strings.HasPrefix(sid, "acoes") {
		if sid == "acoes" {
			getAcoes(w, r, user.ID)
			return
		}
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

// Comprar connection
func comprar(w http.ResponseWriter, r *http.Request) {
	compra := queryGetCompraAcoes()
	json, _ := json.Marshal(compra)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func getAcoes(w http.ResponseWriter, r *http.Request, id int) {
	acoes := queryGetAcoesUsuario(id)
	json, _ := json.Marshal(acoes)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
func comprasLances(w http.ResponseWriter, r *http.Request, codigo string) {
	lances := queryGetLancesAcao(codigo)
	json, _ := json.Marshal(lances)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func venderAcao(w http.ResponseWriter, r *http.Request, valor float64, codigo, user string, qtd int) {
	err := insertCompra(user, codigo, valor, qtd)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// GraficH Gerencia todos os graficos
func GraficH(w http.ResponseWriter, r *http.Request) {
	auth, _ := checkUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sid := strings.TrimPrefix(r.URL.Path, "/grafic/")
	println(sid)

	w.Header().Set("Content-Type", "application/json")
	if sid == "saldo" {
		saldo(w, r)
	}
}

// UserH send user info or get per user
func UserH(w http.ResponseWriter, r *http.Request) {

	auth, u := checkUser(r)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sid := strings.TrimPrefix(r.URL.Path, "/user/")
	println(sid)

	w.Header().Set("Content-Type", "application/json")
	if sid == "info" {
		json, _ := json.Marshal(u)
		fmt.Fprint(w, string(json))
	}
}

// NewSesseion try create new session
func NewSesseion(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	erro := !checkCredentials(*creds)

	if erro {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// Signup create new user
func Signup(w http.ResponseWriter, r *http.Request) {

	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = createUser(*creds)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func checkUser(r *http.Request) (valid bool, userOb User) {
	user, pass, _ := r.BasicAuth()

	cred := Credentials{
		Username: user,
		Password: pass,
	}

	valid = checkCredentials(cred)

	if valid {
		id := queryGetUseID(cred.Username)
		userOb = User{
			Username: user,
			ID:       id,
		}
	}
	return
}

func checkCredentials(cred Credentials) bool {

	password, err := queryGetPassword(cred)

	if err != nil {
		fmt.Println("login:" + cred.Username)
		fmt.Println(err.Error())
		return false
	}

	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(cred.Password)); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
