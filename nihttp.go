package nihttp

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	http.Client

	header map[string]string
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		http.Client{
			Jar:     jar,
			Timeout: time.Duration(10) * time.Second,
		},
		map[string]string{},
	}
}

func (c *Client) Insecure() *Client {
	c.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return c
}

func (c *Client) AddHeader(key, value string) *Client {
	c.header[key] = value
	return c
}

func (c *Client) ClearHeader() *Client {
	c.header = map[string]string{}
	return c
}

func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range c.header {
		req.Header.Set(k, v)
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetJson(url string, out interface{}) error {
	client := NewClient()

	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	DecodeJson(resp, out)
	return nil
}

func DecodeJson(resp *http.Response, out interface{}) error {
	return json.NewDecoder(resp.Body).Decode(out)
}
