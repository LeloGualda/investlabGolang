package alphavantage

import (
	"testing"

	"../querys"
)

func TestCreateBot(t *testing.T) {
	querys.Connection("root:123456@/mydb")
	AddAcaoBot("BOVA11.SAO")
}
