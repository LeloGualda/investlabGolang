package server

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"../alphavantage"
	"../handles"
	"../querys"
)

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("entrou no log")
		h.ServeHTTP(w, r)
	})
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	staticRouter := "./bin/web/static"
	// s := strings.sta(r.URL.Path, "/admin")
	// _, user := handles.CheckUser(r)

	// fmt.Println("USER NAME")
	// fmt.Println(r.URL.Path)
	// fmt.Println(user.ID)

	if strings.HasPrefix(r.URL.Path, "/admin") {
		lp := filepath.Join(staticRouter, "admin.html")

		tmpl, err := template.ParseFiles(lp)
		if err != nil {
			fmt.Println("46:" + err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "admin.html", nil); err != nil {
			fmt.Println("54:" + err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	} else {

		lp := filepath.Join(staticRouter, "index.html")

		tmpl, err := template.ParseFiles(lp)
		if err != nil {
			fmt.Println("46:" + err.Error())
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
			fmt.Println("54:" + err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	}

}

// Start server
func Start() {

	mysqlCon := flag.String("mysqlCon", "root:123456@/mydb", "set mysqlCon, user:password@/dbname")
	flag.Parse()

	paths := "bin/web/public"

	fmt.Println("Mysql Connection...")
	querys.Connection(*mysqlCon)

	fmt.Println("Starting server...")
	alphavantage.CreateUserBoot()
	alphavantage.CreateUserAdmin()
	http.HandleFunc("/session", handles.NewSesseion)
	http.HandleFunc("/test", handles.Test)
	http.HandleFunc("/api/", handles.Api)
	http.HandleFunc("/signup", handles.Signup)
	http.HandleFunc("/user", handles.UserH)
	http.HandleFunc("/user/info", handles.UserH)
	fs := http.FileServer(http.Dir(paths))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", serveTemplate)

	fmt.Printf("run in http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
