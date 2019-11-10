package handles

import "testing"

func TestDrop(t *testing.T) {
	Connection("root:123456@/mydb")
	dropDb()
}
