package api

import (
	"io/ioutil"
	"net/http"
)

// Client is a basic wrapper around the default
// net/http client.
type Client struct {
	http.Client
}

// Get performs a GET request on the URL provided
func (c *Client) Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}