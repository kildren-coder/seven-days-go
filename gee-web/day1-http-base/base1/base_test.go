package base

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttp(t *testing.T) {
	t.Run("return / response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		indexHandler(response, request)

		got := response.Body.String()
		want := "URL.Path = \"/\"\n"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("return /hello response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		request.Header.Set("Accept", "/")
		request.Header.Set("User-Agent", "curl/7.54.0")

		helloHandler(response, request)

		var build strings.Builder
		for k, v := range request.Header {
			//s1 := string()"Header[%q] = %q\n", k, v
			build.WriteString("Header[")
			build.WriteString(k)
			build.WriteString("] = [")
			build.WriteString(v[0] + "]")
			build.WriteString("\n")
		}

		//fmt.Println(build.String() + "23")

		got := response.Body.String()
		want := build.String()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})
}
