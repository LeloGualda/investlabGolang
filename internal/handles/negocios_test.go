package handles

// import (
// 	"fmt"
// 	"testing"
// )

// func TestInsertDinheiro(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	insertDinheiro("MSFT.MEX", 200.50)
// }

// func TestInsertDinheiroError(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	insertDinheiro("OUTRO USER", 200.50)
// }

// func TestInsertDinheiroProf(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	insertDinheiro("prof", 1000.00)
// }
// func TestInsertTransCompra(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	trans := queryGetTransCompra(1)
// 	println(fmt.Sprintf("%f", trans.valorVenda))
// }

// func TestGetAcoes(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	compras := queryGetCompraAcoes()
// 	for _, el := range compras {
// 		fmt.Println("entrou")
// 		fmt.Println(el.Codigo)
// 		fmt.Println(el.Date)
// 	}
// }

// func TestLancesAcoes(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	compras := queryGetLancesAcao("MSFT.MEX")
// 	for _, el := range compras {
// 		fmt.Println("entrou")
// 		fmt.Println(el.Codigo)
// 		println(fmt.Sprintf("%f", el.Valor))
// 	}
// }

// func TestRemoveAcao(t *testing.T) {
// 	Connection("root:123456@/mydb")
// 	removeCarteira(1)
// }
