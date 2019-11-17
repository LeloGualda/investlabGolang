package handles

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func checkUser(r *http.Request) (valid bool, userOb User) {
	user, pass, _ := r.BasicAuth()

	cred := Credentials{
		Username: user,
		Password: pass,
	}

	valid = checkCredentials(cred)

	if valid {
		id := queryGetUseID(cred.Username)
		userOb = User{
			Username: user,
			ID:       id,
		}
	}
	return
}

func checkCredentials(cred Credentials) bool {

	password, err := queryGetPassword(cred)

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
