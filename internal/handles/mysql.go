package handles

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// type UserInfo struct {
// 	Name    string `json:"name", db:"name"`
// 	Email   string `json:"email", db:"email"`
// 	Id_user int    `json:"id_user", db:"_user"`
// }

type User struct {
	Username string `json:"username", db:"username"`
	ID       int    `json:"id", db:"id_user"`
}
type Carteira struct {
	ID     int    `json:"id", db:"id_user"`
	Codigo string `json:"codigo", db:"codigo"`
	Venda  bool   `json:"venda", db:"venda"`
	Lance  int    `json:"lance", db:"lance"`
}

type TransacaoCompra struct {
	Codigo      string  `json:"codigo", db:"codigo"`
	Data        string  `json:"data", db:"data"`
	Valor       float64 `json:"valor", db:"valor"`
	valorVenda  float64 `json:"valorVenda", db:"valorVenda"`
	Lance       int     `json:"lance", db:"lance"`
	IdUserVenda int     `json:"id_user", db:"id_user"`
}

type Lances struct {
	Codigo string  `json:"codigo", db:"codigo"`
	Valor  float64 `json:"valor", db:"valor"`
	Qtd    int     `json:"qtd", db:"qtd"`
}

type CarteiraUser struct {
	IDUser     int `json:"iduser", db:"id_user"`
	IDCarteira int `json:"idcarteira", db:"id_carteira"`
}

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

func queryGetUseID(username string) int {
	u := &User{}
	err := db.QueryRow("SELECT id_user FROM usuario WHERE username = ?", username).Scan(&u.ID)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return u.ID
}

func queryGetPassword(cred Credentials) (password string, err error) {
	storedCreds := &Credentials{}
	query := "select password from usuario where username='" + cred.Username + "'"
	result := db.QueryRow(query)
	err = result.Scan(&storedCreds.Password)
	password = storedCreds.Password
	return
}

// func queryGetAcoes() []ApiAcao {
// 	var apis = []ApiAcao{}

// 	query := "select * from acoes"
// 	rows, _ := db.Query(query)

// 	for rows.Next() {
// 		api := &ApiAcao{}
// 		rows.Scan(&api.Codigo, &api.Name, &api.Region, &api.Type, &api.Currency)
// 		apis = append(apis, *api)
// 	}

// 	return apis
// }

// func insertAcao(acao ApiAcao) {
// 	query := "INSERT INTO acoes (codigo,nome,regiao,tipo,moeda) VALUES (?,?,?,?,?);"
// 	_, err := db.Query(query, acao.Codigo, acao.Name, acao.Region, acao.Type, acao.Currency)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

func insertCarteira(carteira Carteira) {
	query := "INSERT INTO carteira (id_user,codigo,venda,lance) VALUES (?,?,?,?);"
	_, err := db.Query(query, carteira.ID, carteira.Codigo, carteira.Venda, carteira.Lance)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// func insertValorAcao(valor BaseValores) {
// 	query := "INSERT INTO valores (codigo,data,valor) VALUES (?,?,?);"

// 	_, err := db.Query(query, valor.Codigo, valor.Date, valor.Valor)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

func insertTransacao(trans Transacao) {
	query := "INSERT INTO transacao(data,valor,id_user,tipo,descricao) VALUES (?,?,?,?,?);"

	date := ""

	if trans.Data == "" {
		date = "2019-09-08"
	} else {
		date = trans.Data
	}
	_, err := db.Query(query, date, trans.Valor, trans.ID, trans.Tipo, trans.Descricao)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createUser(creds Credentials) error {
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

func queryGetTransCompra(idAcao int) TransacaoCompra {
	transC := &TransacaoCompra{}
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

	err := result.Scan(&transC.Data, &transC.Lance, &transC.Lance, &transC.Codigo, &transC.IdUserVenda, &transC.valorVenda)

	if err != nil {
		fmt.Println(err.Error())
	}

	return *transC
}

func removeCarteira(id int) {
	db.Query("DELETE FROM carteira WHERE carteira.id_carteira = ?;", id)

}

func queryGetLancesAcao(codigo string, id_user int) []Lances {

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

	var lances = []Lances{}

	rows, err := db.Query(query, id_user)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		lance := &Lances{}

		rows.Scan(&lance.Codigo, &lance.Qtd, &lance.Valor)

		lances = append(lances, *lance)
	}
	return lances
}

func queryGetCarteiraUser(codigo string, valor float64) []CarteiraUser {

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

	var carts = []CarteiraUser{}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		car := &CarteiraUser{}

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
