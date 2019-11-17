package handles

import (
	"fmt"
	"testing"
)

func TestGetSaldo(t *testing.T) {
	Connection("root:123456@/mydb")
	q := queryGetSaldo(1)
	fmt.Printf("%f", q.Valor)
}
