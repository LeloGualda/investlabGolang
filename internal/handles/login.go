package handles

import (
	"fmt"
	"net/http"

	"../querys"
	"../structs"
	"golang.org/x/crypto/bcrypt"
)

func CheckUser(r *http.Request) (valid bool, userOb structs.User) {
	user, pass, _ := r.BasicAuth()

	cred := structs.Credentials{
		Username: user,
		Password: pass,
	}

	valid = checkCredentials(cred)

	if valid {
		id := querys.QueryGetUseID(cred.Username)
		userOb = structs.User{
			Username: user,
			ID:       id,
		}
	}
	return
}

func checkCredentials(cred structs.Credentials) bool {

	password, err := querys.QueryGetPassword(cred)

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
