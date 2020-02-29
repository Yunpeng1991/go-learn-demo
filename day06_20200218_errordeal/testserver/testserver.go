package testserver

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const httpPath = "/test/"

func TestServer() {
	http.HandleFunc("/test/", func(writer http.ResponseWriter, request *http.Request) {
		if strings.Index(request.URL.Path, httpPath) != 0 {
			http.Error(writer, fmt.Sprintf("path not exist %s", request.URL.Path), http.StatusNotFound)
			return
		}
		path := request.URL.Path[len(httpPath):]
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
	log.Warn("http error ", err.Error())
	code := http.StatusOK
	switch {
	case os.IsNotExist(err):
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	http.Error(writer, http.StatusText(code), code)
}
