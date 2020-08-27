// Package Requestor contains the methods to make HTTP requests to different endpoints
package Requestor

import "net/http"

// Client here is a struct that holds different configurations to tune Requestor
type Client struct {
	MaxRetriesOnError int8
}

// Get performs a HTTP GET request. It takes in a URL, user specified headers, query params and returns Response and
// error if exist
func (c *Client) Get(url string, headers, queryParams map[string][]string) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodGet, headers, queryParams, nil)
}

// Head performs a HTTP HEAD request. It takes in a URL, user specified headers, query params and returns Response and
// error if exist
func (c *Client) Head(url string, headers, queryParams map[string][]string) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodHead, headers, queryParams, nil)
}

// Post performs a HTTP POST request. It takes in a URL, user specified headers, query params, data and returns
// Response and error if exist
func (c *Client) Post(url string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodPost, headers, queryParams, data)
}

// Put performs a HTTP PUT request. It takes in a URL, user specified headers, query params, data and returns
// Response and error if exist
func (c *Client) Put(url string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodPut, headers, queryParams, data)
}

// Patch performs a HTTP PATCH request. It takes in a URL, user specified headers, query params, data and returns
// Response and error if exist
func (c *Client) Patch(url string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodPatch, headers, queryParams, data)
}

// Delete performs a HTTP DELETE request. It takes in a URL, user specified headers, query params, data and returns
// Response and error if exist
func (c *Client) Delete(url string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodDelete, headers, queryParams, data)
}

// Connect performs a HTTP CONNECT request. It takes in a URL, user specified headers, query params and returns
// Response error if exist
func (c *Client) Connect(url string, headers, queryParams map[string][]string) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodConnect, headers, queryParams, nil)
}

// Options performs a HTTP Options request. It takes in a URL, user specified headers, query params and returns
// Response and error if exist
func (c *Client) Options(url string, headers, queryParams map[string][]string) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodOptions, headers, queryParams, nil)
}

// Trace performs a HTTP TRACE request. It takes in a URL, user specified headers, query params and returns Response
// and error if exist
func (c *Client) Trace(url string, headers, queryParams map[string][]string) (response *http.Response, err error) {

	return c.makeRequest(url, http.MethodTrace, headers, queryParams, nil)
}

// Custom performs a HTTP request with custom method that the server accepts. It takes in a URL, custom method,
// user specified headers, query params and returns Response and error if exist
func (c *Client) Custom(url, method string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {

	return c.makeRequest(url, method, headers, queryParams, data)
}

// makeRequest is a helper method for the above HTTP methods
func (c *Client) makeRequest(url, method string, headers map[string][]string, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {

	return response, nil
}