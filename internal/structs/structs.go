package structs

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
	ValorVenda  float64 `json:"valorVenda", db:"valorVenda"`
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
type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

type AcoesDoUsuario struct {
	Codigo string `json:"codigo", db:"codigo"`
	Lance  int    `json:"lance", db:"lance"`
	Qtd    int    `json:"qtd", db:"qtd"`
	Venda  bool   `json:"venda", db:"venda"`
}
type Transacao struct {
	ID        int     `json:"id", db:"id_transacao"`
	Data      string  `json:"data", db:"data"`
	Valor     float64 `json:"valor", db:"valor"`
	Tipo      int     `json:"tipo", db:"tipo"`
	Descricao string  `json:"descricao", db:"descricao"`
}

type ApiAcao struct {
	Codigo   string `json:"1. symbol", db:"codigo"`
	Name     string `json:"2. name", db:"nome"`
	Type     string `json:"3. type", db:"tipo"`
	Region   string `json:"4. region", db:"regiao"`
	Currency string `json:"4. region", db:"moeda"`
}

type ApiPesquisa struct {
	BestMatches []ApiAcao `json:bestMatches`
}

type ApiValorAcao struct {
	Open string `json:"1. open"`
}
type ApiInfoAcao struct {
	Codigo string `json:"2. Symbol"`
}

type ApiSeriesDay struct {
	Time map[string]ApiValorAcao `json:"Time Series (Daily)"`
	Info ApiInfoAcao             `json:"Meta Data"`
}

type BaseValores struct {
	Codigo string  `json:"codigo",db:"codigo"`
	Date   string  `json:"data",db:"data"`
	Valor  float64 `json:"valor",db:"valor"`
}
