package requestor

import (
	"net/http"
	"testing"
)

func TestClient_SetHTTPProxy(t *testing.T) {
	testCases := []struct {
		proxyURL string
		username string
		password string
	}{
		{"localhost:1234", "", ""},
		{"localhost:1234", "root", "toor"},
	}

	for _, testCase := range testCases {
		client := New()
		client.SetHTTPProxy(testCase.proxyURL, testCase.username, testCase.password)

		url, err := client.transport.Proxy(&http.Request{})
		if err != nil {
			t.Error(err)
		}

		if url.Host != testCase.proxyURL {
			t.Errorf("Expected: %s \n Got: %s", testCase.proxyURL, url.Host)
		}
	}
}

func TestClient_SetHTTPSProxy(t *testing.T) {
	testCases := []struct {
		proxyURL string
		username string
		password string
	}{
		{"localhost:1234", "", ""},
		{"localhost:1234", "root", "toor"},
	}

	for _, testCase := range testCases {
		client := New()
		client.SetHTTPSProxy(testCase.proxyURL, testCase.username, testCase.password)

		url, err := client.transport.Proxy(&http.Request{})
		if err != nil {
			t.Error(err)
		}

		if url.Host != testCase.proxyURL {
			t.Errorf("Expected: %s \n Got: %s", testCase.proxyURL, url.Host)
		}
	}
}
