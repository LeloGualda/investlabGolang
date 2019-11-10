package handles

import (
	"database/sql"
	"fmt"
	"strconv"

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

type Transacao struct {
	ID        int     `json:"id", db:"id_transacao"`
	Data      string  `json:"data", db:"data"`
	Valor     float64 `json:"valor", db:"valor"`
	Tipo      int     `json:"tipo", db:"tipo"`
	Descricao string  `json:"descricao", db:"descricao"`
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

type AcoesDoUsuario struct {
	Codigo string `json:"codigo", db:"codigo"`
	Lance  int    `json:"lance", db:"lance"`
	Qtd    int    `json:"qtd", db:"qtd"`
	Venda  bool   `json:"venda", db:"venda"`
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
	err := db.QueryRow("SELECT id_user FROM mydb.usuario WHERE username = ?", username).Scan(&u.ID)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return u.ID
}

func queryGetSaldo(id_user int) float64 {
	u := &Transacao{}
	err := db.QueryRow("select sum(valor) from mydb.transacao where id_user = ?", id_user).Scan(&u.Valor)
	if err != nil {
		fmt.Println(err.Error())
	}
	return u.Valor
}

func queryGetPassword(cred Credentials) (password string, err error) {
	storedCreds := &Credentials{}
	query := "select password from usuario where username='" + cred.Username + "'"
	result := db.QueryRow(query)
	err = result.Scan(&storedCreds.Password)
	password = storedCreds.Password
	return
}

func queryGetAcoes() []ApiAcao {
	var apis = []ApiAcao{}

	query := "select * from acoes"
	rows, _ := db.Query(query)

	for rows.Next() {
		api := &ApiAcao{}
		rows.Scan(&api.Codigo, &api.Name, &api.Region, &api.Type, &api.Currency)
		apis = append(apis, *api)
	}

	return apis
}

func insertAcao(acao ApiAcao) {
	query := "INSERT INTO mydb.acoes (codigo,nome,regiao,tipo,moeda) VALUES (?,?,?,?,?);"
	_, err := db.Query(query, acao.Codigo, acao.Name, acao.Region, acao.Type, acao.Currency)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func insertCarteira(carteira Carteira) {
	query := "INSERT INTO mydb.carteira (id_user,codigo,venda,lance) VALUES (?,?,?,?);"
	_, err := db.Query(query, carteira.ID, carteira.Codigo, carteira.Venda, carteira.Lance)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func insertValorAcao(valor BaseValores) {
	query := "INSERT INTO mydb.valores (codigo,data,valor) VALUES (?,?,?);"

	_, err := db.Query(query, valor.Codigo, valor.Date, valor.Valor)
	if err != nil {
		fmt.Println(err.Error())
	}
}

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
	query := "insert into mydb.usuario (username,password) values ('" + creds.Username + "','" + string(hashedPassword) + "')"
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

	select mydb.valores.data,
		   mydb.valores.valor,
		   mydb.carteira.lance,
		   mydb.valores.codigo,
		   mydb.carteira.id_user,
		   mydb.valores.valor + (mydb.valores.valor *  (mydb.carteira.lance/100)) as valorVenda

		from mydb.valores,mydb.carteira
			where mydb.valores.codigo =  mydb.carteira.codigo
			and
			data = (select max(data) from mydb.valores)
			and
			carteira.id_carteira = ?`

	result := db.QueryRow(query, idAcao)

	err := result.Scan(&transC.Data, &transC.Lance, &transC.Lance, &transC.Codigo, &transC.IdUserVenda, &transC.valorVenda)

	if err != nil {
		fmt.Println(err.Error())
	}

	return *transC
}

func queryGetCompraAcoes() []BaseValores {

	query := `
	select
	data,
	valor,
	codigo
	from mydb.valores
		where
	data = (select max(data) from mydb.valores group by codigo);
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

func removeCarteira(id int) {
	db.Query("DELETE FROM mydb.carteira WHERE mydb.carteira.id_carteira = ?;", id)

}

func queryGetLancesAcao(codigo string) []Lances {

	query := `
	select

	codigo,
	count(lance) as qtd,
	valor + (valor *  (lance/100)) as valor
	
	from (select  
		   mydb.carteira.codigo,
		   mydb.carteira.lance,
		   mydb.valores.valor
		   from mydb.carteira,mydb.valores
			   where 
				   venda = true 
				   and 
				  mydb.valores.codigo = mydb.carteira.codigo
						   and 
				   mydb.valores.data = (select max(data) from mydb.valores group by codigo)) as a
				   where codigo = "` + codigo + `"group by codigo,lance,valor order by lance, qtd desc;`

	var lances = []Lances{}

	rows, err := db.Query(query)
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
		mydb.carteira.id_carteira,
        mydb.carteira.id_user
		from mydb.carteira,mydb.valores
			where
				venda = true
                and
				valor + (valor *  (lance/100)) = "` + fmt.Sprintf("%f", valor) + `"
			   and
			   mydb.carteira.codigo = '` + codigo + `'
		       and
				mydb.valores.data = (select max(data) from mydb.valores group by codigo);
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

// SELECT id_user,lance,codigo FROM mydb.carteira where  mydb.carteira.id_carteira  = 1;
// select data,valor,codigo
// from mydb.valores
// 	where data = (select max(data) from mydb.valores group by codigo);

// select data,valor,codigo
// from mydb.valores
// 	where
// 		data = (select max(data) from mydb.valores)
// 		and
// 		codigo = (SELECT codigo FROM mydb.carteira where  mydb.carteira.id_carteira  = 1);
