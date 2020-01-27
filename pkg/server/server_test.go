package server_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/mtebourbi/lbc-fizzbuzz/pkg/fizzbuzz"
	"github.com/mtebourbi/lbc-fizzbuzz/pkg/server"
)

func TestFizzBuzzHandler(t *testing.T) {
	ts := httptest.NewServer(server.RegisterRoutes())
	defer ts.Close()

	url := "/fizzbuzz?int1=3&int2=5&limit=30&str1=fizz&str2=buzz"
	res, err := http.Get(ts.URL + url)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("http response code: %v", res.StatusCode)
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("request response body error: %v", err)
	}

	fz, _ := fizzbuzz.FizzBuzz(3, 5, 30, "fizz", "buzz")
	expect := strings.Join(fz, ",") + "\n"

	if s := string(out); s != expect {
		t.Errorf("Expected %v found : %v", expect, s)
	}
}

func TestFizzBuzzHandlerBadRequest(t *testing.T) {
	urls := []struct {
		name string
		url  string
	}{
		{
			name: "invalid int1",
			url:  "/fizzbuzz?int1=xx&int2=5&limit=30&str1=fizz&str2=buzz",
		},
		{
			name: "invalid int2",
			url:  "/fizzbuzz?int1=1&int2=xx&limit=30&str1=fizz&str2=buzz",
		},
		{
			name: "invalid int2",
			url:  "/fizzbuzz?int1=1&int2=5&limit=xx&str1=fizz&str2=buzz",
		},
	}

	ts := httptest.NewServer(server.RegisterRoutes())
	defer ts.Close()

	for _, u := range urls {
		ft := func(t *testing.T) {
			res, err := http.Get(ts.URL + u.url)
			if err != nil {
				t.Fatalf("request error: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusBadRequest {
				t.Errorf("http response code: %v", res.StatusCode)
			}
		}
		t.Run(u.name, ft)
	}

}

func TestTopRequestHandler(t *testing.T) {
	ts := httptest.NewServer(server.RegisterRoutes())
	defer ts.Close()

	server.DefaultRequestRepo = server.NewRequestRepo()
	var wg sync.WaitGroup
	wg.Add(2)

	doReq := func(url string, count int) {
		defer wg.Done()
		for ; count > 0; count-- {
			res, err := http.Get(ts.URL + url)
			if err != nil {
				t.Fatalf("request error: %v", err)
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Fatalf("http response code: %v", res.StatusCode)
			}

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("request response body error: %v", err)
			}
		}
	}

	go doReq("/fizzbuzz?int1=3&int2=5&limit=30&str1=fizz&str2=buzz", 100)
	go doReq("/fizzbuzz?int1=3&int2=5&limit=30&str1=toto&str2=tata", 60)

	wg.Wait()

	url := "/tophits"

	res, err := http.Get(ts.URL + url)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("http response code: %v", res.StatusCode)
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("request response body error: %v", err)
	}

	var expect = `{"int1":3,"int2":5,"limit":30,"str1":"fizz","str2":"buzz","hits":100}`

	if s := string(out); s != expect {
		t.Errorf("Expected %v found : %v", expect, s)
	}

}

func TestTopRequestHandlerInitial(t *testing.T) {
	url := "/tophits"

	ts := httptest.NewServer(server.RegisterRoutes())
	defer ts.Close()

	server.DefaultRequestRepo = server.NewRequestRepo()

	res, err := http.Get(ts.URL + url)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("http response code: %v", res.StatusCode)
	}

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("request response body error: %v", err)
	}

	var expect = `{"int1":0,"int2":0,"limit":0,"str1":"","str2":"","hits":0}`

	if s := string(out); s != expect {
		t.Errorf("Expected %v found : %v", expect, s)
	}

}

func TestHealthCheckHandler(t *testing.T) {
	url := "/ping"

	ts := httptest.NewServer(server.RegisterRoutes())
	defer ts.Close()

	res, err := http.Get(ts.URL + url)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("http response code: %v", res.StatusCode)
	}

}
