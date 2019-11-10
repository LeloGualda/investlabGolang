package provider

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Upload make upload file
func Upload(w http.ResponseWriter, r *http.Request) {

	var diretoryFiles = "/bin/files/"

	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("error get bytes file")
		fmt.Println(err)
	}

	filename := getPathRunner() + diretoryFiles + handler.Filename

	println("filename: " + filename)
	println("Exists file: " + strconv.FormatBool(Exists(filename)))

	if Exists(filename) {
		errWrite := ioutil.WriteFile(filename, fileBytes, 0644)
		isError(errWrite)
	} else {
		createFile(filename, fileBytes)
	}
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createFile(path string, content []byte) bool {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return false
		}
		file.Write(content)
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
	return true
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func getPathRunner() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
