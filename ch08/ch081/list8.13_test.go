package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestApi8_13(t *testing.T) {
	tt := func(t *testing.T) {
		ct := "application/vnd.mytools.json; version=2.0"
		req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
		req.Header.Set("Accept", ct)
		res, _ := http.DefaultClient.Do(req)
		if res.Header.Get("Content-Type") != ct {
			t.Error("Unexpected content type returned")
			return
		}
		b, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Printf("%s\n", b)
	}

	t.Run("test api", tt)
}

func TestApi8_13_1(t *testing.T) {
	ct := "application/vnd.mytools.json; version=2.0"
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req.Header.Set("Accept", ct)
	res, _ := http.DefaultClient.Do(req)
	if res.Header.Get("Content-Type") != ct {
		t.Error("Unexpected content type returned")
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", b)
}
