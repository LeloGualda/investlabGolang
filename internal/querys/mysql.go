package querys

import (
	"database/sql"
	"fmt"

	"../structs"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// type UserInfo struct {
// 	Name    string `json:"name", db:"name"`
// 	Email   string `json:"email", db:"email"`
// 	Id_user int    `json:"id_user", db:"_user"`
// }

func Connection(mysqlCon string) {
	var err error
	fmt.Println(mysqlCon)
	db, err = sql.Open("mysql", mysqlCon)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("conected mysql")
}

func QueryGetUseID(username string) int {
	u := &structs.User{}
	err := db.QueryRow("SELECT id_user FROM usuario WHERE username = ?", username).Scan(&u.ID)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return u.ID
}

func QueryGetPassword(cred structs.Credentials) (password string, err error) {
	storedCreds := &structs.Credentials{}
	query := "select password from usuario where username='" + cred.Username + "'"
	result := db.QueryRow(query)
	err = result.Scan(&storedCreds.Password)
	password = storedCreds.Password
	return
}

func QueryGetAcoes() []structs.ApiAcao {
	var apis = []structs.ApiAcao{}

	query := "select * from acoes"
	rows, _ := db.Query(query)

	for rows.Next() {
		api := &structs.ApiAcao{}
		rows.Scan(&api.Codigo, &api.Name, &api.Region, &api.Type, &api.Currency)
		apis = append(apis, *api)
	}

	return apis
}

func QueryGetAcoesEspecifica(codigo string) []structs.ApiAcao {
	var apis = []structs.ApiAcao{}

	query := "select * from acoes where codigo = ?"
	rows, _ := db.Query(query, codigo)

	for rows.Next() {
		api := &structs.ApiAcao{}
		rows.Scan(&api.Codigo, &api.Name, &api.Region, &api.Type, &api.Currency)
		apis = append(apis, *api)
	}

	return apis
}

func InsertTransacao(trans structs.Transacao) {
	query := "INSERT INTO transacao(data,valor,id_user,tipo,descricao) VALUES (?,?,?,?,?);"

	date := ""

	if trans.Data == "" {
		date = newDate()
	} else {
		date = trans.Data
	}
	_, err := db.Query(query, date, trans.Valor, trans.ID, trans.Tipo, trans.Descricao)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateUser(creds structs.Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	query := "insert into usuario (username,password) values ('" + creds.Username + "','" + string(hashedPassword) + "')"
	_, err = db.Query(query)
	return err
}

func dropDb() {
	_, err := db.Exec("DROP DATABASE IF EXISTS mydb;")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func QueryGetTransCompra(idAcao int) structs.TransacaoCompra {
	transC := &structs.TransacaoCompra{}

	query := `

	select valores.data,
		   valores.valor,
		   carteira.lance,
		   valores.codigo,
		   carteira.id_user,
		   valores.valor + (valores.valor *  (carteira.lance/100)) as valorVenda

		from valores,carteira
			where valores.codigo =  carteira.codigo
			and
			data = (select max(data) from valores)
			and
			carteira.id_carteira = ?`

	result := db.QueryRow(query, idAcao)

	err := result.Scan(&transC.Data, &transC.Lance, &transC.Lance, &transC.Codigo, &transC.IdUserVenda, &transC.ValorVenda)

	if err != nil {
		fmt.Println(err.Error())
	}

	return *transC
}

func RemoveCarteira(id int) {
	db.Query("DELETE FROM carteira WHERE carteira.id_carteira = ?;", id)

}

func QueryGetLancesAcao(codigo string, id_user int) []structs.Lances {

	query := `
	select

	codigo,
	count(lance) as qtd,
	valor + (valor *  (lance/100)) as valor
	
	from (select
		   carteira.codigo,
		   carteira.lance,
		   valores.valor
		   from carteira,valores
			   where
				   venda = true
				   and
                   carteira.id_user != ?
				   and
				  valores.codigo = carteira.codigo
						   and
				   valores.data = (select max(data) from valores group by codigo)) as a
				   where codigo = "` + codigo + `"group by codigo,lance,valor order by lance, qtd desc;`

	var lances = []structs.Lances{}

	rows, err := db.Query(query, id_user)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		lance := &structs.Lances{}

		rows.Scan(&lance.Codigo, &lance.Qtd, &lance.Valor)

		lances = append(lances, *lance)
	}
	return lances
}

func QueryGetCarteiraUser(codigo string, valor float64) []structs.CarteiraUser {

	query := `
	select  
		carteira.id_carteira,
        carteira.id_user
		from carteira,valores
			where
				venda = true
                and
				valor + (valor *  (lance/100)) = "` + fmt.Sprintf("%f", valor) + `"
			   and
			   carteira.codigo = '` + codigo + `'
		       and
				valores.data = (select max(data) from valores group by codigo);
	`

	var carts = []structs.CarteiraUser{}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		car := &structs.CarteiraUser{}

		rows.Scan(&car.IDCarteira, &car.IDUser)

		carts = append(carts, *car)
	}
	return carts
}

// SELECT id_user,lance,codigo FROM carteira where  carteira.id_carteira  = 1;
// select data,valor,codigo
// from valores
// 	where data = (select max(data) from valores group by codigo);

// select data,valor,codigo
// from valores
// 	where
// 		data = (select max(data) from valores)
// 		and
// 		codigo = (SELECT codigo FROM carteira where  carteira.id_carteira  = 1);
