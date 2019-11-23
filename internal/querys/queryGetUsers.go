package querys

import (
	"fmt"

	"../structs"
)

func QueryGetUsers() []structs.User {

	query := `
	select id_user,username from usuario group by username order by id_user;
	`

	var users = []structs.User{}

	rows, err := db.Query(query)

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		user := &structs.User{}

		rows.Scan(&user.ID, &user.Username)

		users = append(users, *user)
	}
	return users
}
