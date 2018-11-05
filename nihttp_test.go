package nihttp

import (
	"testing"
)

type dataStruct struct {
	Origin  string
	Url     string
	Args    map[string]string
	Headers map[string]string
	Form    map[string]string
}

func TestGetSuccess(t *testing.T) {

	url := "http://httpbin.org/get?x=1"
	client := NewClient()

	resp, _ := client.Get(url)

	var data dataStruct

	err := DecodeJson(resp, &data)
	if err != nil {
		t.Fatal(err)
	}

	if data.Url != url {
		t.Fatal("failed")
	}
}

func TestPostSuccess(t *testing.T) {
	url := "http://httpbin.org/post"
	client := NewClient()

	client.AddValues("x", "1")

	resp, err := client.Post(url)
	if err != nil {
		t.Fatal(err)
	}

	var data dataStruct

	err = DecodeJson(resp, &data)
	if err != nil {
		t.Fatal(err)
	}

	if data.Form["x"] != "1" {
		t.Fatal("failed")
	}
}
