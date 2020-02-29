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
			errorHandler(writer, err)
			//don`t forget return it
			return
		}
		defer open.Close()

		file, err := ioutil.ReadAll(open)
		if err != nil {
			errorHandler(writer, err)
			return
		}
		_, err = writer.Write(file)
		if err != nil {
			errorHandler(writer, err)
			return
		}
	})

	err := http.ListenAndServe(":8808", nil)

	if err != nil {
		panic(err)
	}

}

func errorHandler(writer http.ResponseWriter, err error) {
	code := http.StatusOK
	switch {
	case os.IsNotExist(err):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	http.Error(writer, http.StatusText(code), code)
}
