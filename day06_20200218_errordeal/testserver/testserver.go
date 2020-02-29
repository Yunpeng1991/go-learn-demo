package testserver

import (
	"io/ioutil"
	"net/http"
	"os"
)

func TestServer() {
	http.HandleFunc("/test/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/test/"):]
		open, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer open.Close()

		file, err := ioutil.ReadAll(open)
		if err != nil {
			panic(err)
		}
		writer.Write(file)
	})

	http.ListenAndServe(":8808", nil)

}
