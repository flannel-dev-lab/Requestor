// Package Requestor contains the methods to make HTTP requests to different endpoints
package requestor

import (
	"net/http"
	"net/url"
)

// SetHTTPProxy sets a HTTP proxy to the transport, proxyURL is a required parameter, but username and password
// is optional parameters. Proxy URL should be of format IP:PORT or HOSTNAME:PORT
func (c *Client) SetHTTPProxy(proxyURL, username, password string) {
	proxyConfig := &url.URL{
		Scheme: "http",
		Host:   proxyURL,
	}

	if username != "" && password != "" {
		proxyConfig.User = url.UserPassword(username, password)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyConfig),
	}

	c.transport = transport
}

// SetHTTPSProxy sets a HTTP proxy to the transport, proxyURL is a required parameter, but username and password
// is optional parameters. Proxy URL should be of format IP:PORT or HOSTNAME:PORT
func (c *Client) SetHTTPSProxy(proxyURL, username, password string) {
	proxyConfig := &url.URL{
		Scheme: "https",
		Host:   proxyURL,
	}

	if username != "" && password != "" {
		proxyConfig.User = url.UserPassword(username, password)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyConfig),
	}

	c.transport = transport
}
