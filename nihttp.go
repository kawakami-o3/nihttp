package nihttp

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type Client struct {
	http.Client

	header map[string]string

	values url.Values
}

func NewClient() *Client {
	jar, _ := cookiejar.New(nil)
	return &Client{
		http.Client{
			Jar:     jar,
			Timeout: time.Duration(10) * time.Second,
		},
		map[string]string{},
		url.Values{},
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

func (c *Client) AddValues(key, value string) *Client {
	c.values.Add(key, value)
	return c
}

func (c *Client) ClearValues() *Client {
	c.values = url.Values{}
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

func GetString(url string) (string, error) {
	client := NewClient()

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
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

func (c *Client) Post(url string) (*http.Response, error) {
	resp, err := c.PostForm(url, c.values)
	if err != nil {
		return nil, err
	}

	c.ClearValues()

	return resp, nil
}

func DecodeJson(resp *http.Response, out interface{}) error {
	return json.NewDecoder(resp.Body).Decode(out)
}
