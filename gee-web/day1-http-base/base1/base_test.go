package base

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	t.Run("return / response", func(t *testing.T) {
		request := getNewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		indexHandler(response, request)

		got := response.Body.String()
		want := "URL.Path = /\n"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	/*
	t.Run("return /hello response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		request.Header.Set("Accept", "/")
		request.Header.Set("User-Agent", "curl/7.54.0")

		helloHandler(response, request)

		var build strings.Builder
		build = getHeaderContext(*request, build)

		//fmt.Println(build.String() + "23")

		got := response.Body.String()
		want := build.String()

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})

	*/
}


