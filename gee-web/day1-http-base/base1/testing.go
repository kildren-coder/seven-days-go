package base

import (
	"io"
	"net/http"
	"strings"
)

func getNewRequest(method, url string, body io.Reader) (*http.Request) {
	request, _ := http.NewRequest(method, url, body)
	return request
}

func getHeaderContext(request http.Request, build strings.Builder) strings.Builder {
	for k, v := range request.Header {
		//s1 := string()"Header[%q] = %q\n", k, v
		build.WriteString("Header[")
		build.WriteString(k)
		build.WriteString("] = [")
		build.WriteString(v[0] + "]")
		build.WriteString("\n")
	}

	return build
}
