package Requestor

import "net/http"

type Client struct {
	MaxRetriesOnError int8
}

func (c *Client) Get() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Head() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Post(url string, headers http.Header, queryParams map[string][]string) (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Put() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Patch() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Delete() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Connect() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Options() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Trace() (response *http.Response, err error) {

	return response, nil
}

func (c *Client) Custom() (response *http.Response, err error) {

	return response, nil
}