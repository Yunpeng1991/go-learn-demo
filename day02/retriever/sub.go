package retriever

import (
	"fmt"
	"learn-demo/common"
	"net/http"
	"net/http/httputil"
	"time"
)

type Sub struct {
	Duration time.Duration
}

func (sub *Sub) Get(url string) string {
	isNull := common.IsNull(url)
	if isNull == true {
		err := fmt.Errorf("url cannont be nill")
		return err.Error()
	}
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	response, dumpErr := httputil.DumpResponse(resp, true)
	//close stream
	closeErr := resp.Body.Close()

	if dumpErr != nil || closeErr != nil {
		return err.Error()
	}

	return string(response)
}
