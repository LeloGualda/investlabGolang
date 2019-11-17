package alphavantage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

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

var key = "demo"

func getApiPesquisa(body []byte) (*ApiPesquisa, error) {
	var s = new(ApiPesquisa)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func requestApi(param string) []byte {
	protocol := "https://"
	hostApi := "www.alphavantage.co/"

	fmt.Println(protocol + hostApi)
	req, _ := http.NewRequest("GET", protocol+hostApi+param, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", hostApi)
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		return body
	}
	return nil
}
func getAPi(keyword string, createTimeLine bool) ApiPesquisa {

	param := "query?function=SYMBOL_SEARCH&keywords=" + keyword + "&apikey=" + key
	body := requestApi(param)

	s, _ := getApiPesquisa([]byte(body))

	for _, el := range s.BestMatches {
		// insertAcao(el)
		// time.Sleep(5 * time.Second)
		if createTimeLine {
			getSeriesAPi(el.Codigo)
		}
	}
	return *s
}

func getSeriesAPi(codigo string) {
	param := "query?function=TIME_SERIES_DAILY&symbol=" + codigo + "&apikey=" + key
	body := requestApi(param)
	var s = new(ApiSeriesDay)
	err := json.Unmarshal(body, &s)
	fmt.Println("request feito")
	if err != nil {
		fmt.Println("whoops:", err)
	}
	for k, el := range s.Time {
		valor := new(BaseValores)

		valor.Date = k
		valor.Valor, _ = strconv.ParseFloat(el.Open, 64)
		valor.Codigo = s.Info.Codigo

		// insertValorAcao(*valor)
	}
}

func inserSeriesOfAcoes() {
	// api := queryGetAcoes()
	// for _, el := range api {
	// 	getSeriesAPi(el.Codigo)
	// 	time.Sleep(10 * time.Second)
	// }
}

func createBot() {

	// api := queryGetAcoes()

	// fmt.Println("get all acoes")
	// //getAPi("MSFT", true)

	// for _, el := range api {
	// 	fmt.Println(el.Codigo)
	// 	creds := new(Credentials)
	// 	creds.Password = el.Codigo
	// 	creds.Username = el.Codigo
	// 	err := createUser(*creds)

	// 	fmt.Println("create user")

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	add := randAdd()

	// 	for i := 0; i < add; i++ {
	// 		carteira := new(Carteira)
	// 		carteira.Codigo = el.Codigo
	// 		carteira.ID = queryGetUseID(creds.Username)
	// 		carteira.Lance = randLance()
	// 		carteira.Venda = true
	// 		insertCarteira(*carteira)
	// 	}

	// }
}

func randLance() int {
	return ri(2, 3)
}
func randAdd() int {
	return ri(40, 100)
}
func ri(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
