package nihttp

import (
	"testing"
)

func TestExampleSuccess(t *testing.T) {

	url := "http://httpbin.org/get?x=1"
	client := newClient()

	resp, _ := client.Get(url)

	type dataStruct struct {
		Origin  string
		Url     string
		Args    map[string]string
		Headers map[string]string
	}
	var data dataStruct

	err := DecodeJson(resp, &data)
	if err != nil {
		t.Fatal(err)
	}

	if data.Url != url {
		t.Fatal("failed")
	}
}
