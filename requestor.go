// Package requestor contains the methods to make HTTP requests to different endpoints
package requestor

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Client here is a struct that holds different configurations to tune Requestor
type Client struct {
	// MaxRetriesOnError specifies how many times we should retry when a request to server fails
	MaxRetriesOnError uint8
	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	//
	// A Timeout of zero means no timeout.
	// The Client cancels requests to the underlying Transport
	// as if the Request's Context ended.
	Timeout time.Duration
	// IdleConnectionTimeout specifies how long an idle connection is kept in the connection pool
	// 0 means no timeout
	IdleConnectionTimeout time.Duration
	// MaxConnectionsPerHost specifies the max number of connection per host which include connections in dialing
	// active and idle states. When exceeded request gets cancelled with net/http: request canceled.
	// 0 means no limit
	MaxConnectionsPerHost int
	// MaxIdleConnectionsPerHost specifies the max number of keep-alive connections per host
	MaxIdleConnectionsPerHost int
	// MaxIdleConnections specifies the max number of keep-alive connections
	MaxIdleConnections int
	// DisableKeepAlives makes sure that the transport is used only once by disabling keep-alives
	// Default is true
	DisableKeepAlives bool
	// TLSClientConfig specifies the TLS config to use
	TLSClientConfig *tls.Config

	transport *http.Transport
}

func New() (client *Client) {
	return &Client{
		Timeout:               0,
		DisableKeepAlives:     true,
		IdleConnectionTimeout: 0,
		transport:             &http.Transport{},
	}
}

func (c *Client) SetTLSClientConfig(tlsConfig *tls.Config) {
	c.TLSClientConfig = tlsConfig
}

func (c *Client) DisableKeepAlive(val bool) {
	c.DisableKeepAlives = val
}

func (c *Client) SetMaxConnectionsPerHost(connectionCount int) {
	c.MaxConnectionsPerHost = connectionCount
}

func (c *Client) SetMaxIdleConnectionsPerHost(connectionCount int) {
	c.MaxIdleConnectionsPerHost = connectionCount
}

func (c *Client) SetMaxIdleConnections(connectionCount int) {
	c.MaxIdleConnections = connectionCount
}

func (c *Client) SetMaxRetries(retries uint8) {
	c.MaxRetriesOnError = retries
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func (c *Client) SetIdleConnectionTimeout(timeout time.Duration) {
	c.IdleConnectionTimeout = timeout
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

// Connect performs a HTTP CONNECT request. It takes in a URL, user specified headers, proxyHeaders, query params and returns
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
func (c *Client) makeRequest(url, method string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {
	c.populateTransport()

	// Converting headers to canonical headers
	canonicalHeaders := make(map[string][]string, len(headers))
	for headerKey, headerValue := range headers {
		canonicalHeaders[http.CanonicalHeaderKey(headerKey)] = headerValue
	}

	contentType, ok := canonicalHeaders["Content-Type"]
	if ok && len(contentType) >= 1 {
		if contentType[0] == "application/json" || strings.Contains(contentType[0], "application/json") {
			return c.makeJSONRequest(url, method, headers, queryParams, data)
		}
	}

	return c.makeJSONRequest(url, method, headers, queryParams, data)
}

func (c *Client) makeJSONRequest(url, method string, headers, queryParams map[string][]string, data interface{}) (response *http.Response, err error) {
	var dataBytes []byte
	dataBytes, err = json.Marshal(data)
	if err != nil {
		return response, err
	}

	httpClient := http.Client{
		Transport: c.transport,
		Timeout:   c.Timeout,
	}

	var request *http.Request

	if len(dataBytes) > 0 {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(dataBytes))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return response, err
	}

	q := request.URL.Query()

	for queryKey, queryValues := range queryParams {
		for _, val := range queryValues {
			q.Add(queryKey, val)
		}
	}

	request.URL.RawQuery = q.Encode()

	for headerKey, headerValues := range headers {
		for _, val := range headerValues {
			request.Header.Add(headerKey, val)
		}
	}

	return httpClient.Do(request)
}

// populateTransport populates the transport config
func (c *Client) populateTransport() {
	c.transport.DisableKeepAlives = c.DisableKeepAlives
	c.transport.MaxConnsPerHost = c.MaxConnectionsPerHost
	c.transport.IdleConnTimeout = c.IdleConnectionTimeout
	c.transport.MaxConnsPerHost = c.MaxConnectionsPerHost
	c.transport.MaxIdleConnsPerHost = c.MaxIdleConnectionsPerHost
	c.transport.MaxIdleConns = c.MaxIdleConnections
	c.transport.TLSClientConfig = c.TLSClientConfig
}
