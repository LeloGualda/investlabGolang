package server

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"../handles"
)

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("entrou no log")
		h.ServeHTTP(w, r)
	})
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {

	staticRouter := "./bin/web/static"
	lp := filepath.Join(staticRouter, "index.html")
	fp := filepath.Join(staticRouter, filepath.Clean(r.URL.Path))

	// fmt.Println(lp)
	// fmt.Println(fp)
	// // // Return a 404 if the template doesn't exist
	// _, err := os.Stat(fp)
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Println("31:" + err.Error())
	// 		http.NotFound(w, r)
	// 		return
	// 	}
	// }
	// http.ServeFile(w, r, lp)
	// // Return a 404 if the request is for a directory
	// if info.IsDir() {
	// 	fmt.Println("39:" + err.Error())
	// 	http.NotFound(w, r)
	// 	return
	// }

	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		fmt.Println("46:" + err.Error())
		// Log the detailed error
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Println("response: " + lp)
	fmt.Println("response: " + fp)
	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		fmt.Println("54:" + err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

// Start server
func Start() {

	mysqlCon := flag.String("mysqlCon", "root:123456@/mydb", "set mysqlCon, user:password@/dbname")
	flag.Parse()

	paths := "bin/web/public"

	fmt.Println("Mysql Connection...")

	handles.Connection(*mysqlCon)

	fmt.Println("Starting server...")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Println("entrou na home")
	// 	fmt.Println(r.URL)
	// 	http.ServeFile(w, r, "bin/web/static/index.html")
	// })

	// http.HandleFunc("/comprar/ABEV3", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("entrou na ABEV3")
	// 	fmt.Println(r.URL)
	// 	http.ServeFile(w, r, "bin/web/static/index.html")
	// })

	// user
	http.HandleFunc("/session", handles.NewSesseion)
	http.HandleFunc("/test", handles.Test)
	http.HandleFunc("/api/", handles.Api)
	http.HandleFunc("/signup", handles.Signup)
	http.HandleFunc("/user", handles.UserH)
	http.HandleFunc("/user/info", handles.UserH)
	// http.HandleFunc("/logger", logger.Info)
	// http.HandleFunc("/upload", provider.Upload)
	fs := http.FileServer(http.Dir(paths))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", serveTemplate)

	// http.Handle("/public/", http.StripPrefix("bin/web/static/index.html", fs))

	//	http.HandleFunc("/", serveTemplate)

	fmt.Printf("run in http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
