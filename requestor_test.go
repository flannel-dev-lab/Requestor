package requestor

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestServerResponse struct {
	Args     map[string][]string `json:"args"`
	Headers  http.Header         `json:"headers"`
	FormData map[string][]string `json:"form_data"`
	Data     interface{}         `json:"data"`
}

func TestClientOptions(t *testing.T) {
	client := New()

	client.SetTLSClientConfig(&tls.Config{})

	client.DisableKeepAlive(true)
	if !client.DisableKeepAlives {
		t.Errorf("Expected: %t \n Got: %t", true, client.DisableKeepAlives)
	}

	client.SetMaxConnectionsPerHost(1)
	if client.MaxConnectionsPerHost != 1 {
		t.Errorf("Expected: %d \n Got: %d", 1, client.MaxConnectionsPerHost)
	}

	client.SetMaxIdleConnectionsPerHost(1)
	if client.MaxIdleConnectionsPerHost != 1 {
		t.Errorf("Expected: %d \n Got: %d", 1, client.MaxIdleConnectionsPerHost)
	}

	client.SetIdleConnectionTimeout(1)
	if client.IdleConnectionTimeout != 1 {
		t.Errorf("Expected: %d \n Got: %d", 1, client.IdleConnectionTimeout)
	}

	client.SetMaxIdleConnections(1)
	if client.MaxIdleConnections != 1 {
		t.Errorf("Expected: %d \n Got: %d", 1, client.MaxIdleConnections)
	}

	client.SetTimeout(1)
	if client.Timeout != 1 {
		t.Errorf("Expected: %d \n Got: %d", 1, client.Timeout)
	}

	client.SetMaxRetries(2, 1)
	if client.MaxRetriesOnError != 2 {
		t.Errorf("Expected: %d \n Got: %d", 2, client.MaxRetriesOnError)
	}

	client.SetMaxRetries(2, 0)
	if client.MaxRetriesOnError != 2 {
		t.Errorf("Expected: %d \n Got: %d", 2, client.MaxRetriesOnError)
	}
}

func TestClient_Get(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Get(testServer.URL, headers, queryParams)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}
}

func TestClient_Trace(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodTrace {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}))

	client := New()
	resp, err := client.Trace(testServer.URL, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d \n Got: %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_Connect(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodConnect {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}))

	client := New()
	resp, err := client.Connect(testServer.URL, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d \n Got: %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_Options(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodOptions {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}))

	client := New()
	resp, err := client.Options(testServer.URL, nil, nil)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d \n Got: %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_Head(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodHead {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}))

	client := New()
	resp, err := client.Head(testServer.URL)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d \n Got: %d", http.StatusOK, resp.StatusCode)
	}
}

func TestClient_Post(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()
		response.Data = map[string]string{"hello": "world"}

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/json"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Post(testServer.URL, headers, queryParams, map[string]string{"hello": "world"})
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Data.(map[string]interface{})["hello"] != "world" {
		t.Errorf("Expected: %s \n Got: %s", "world", testServerResp.Data.(map[string]string)["hello"])
	}
}

func TestClient_InvalidURL(t *testing.T) {
	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/json"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	_, err := client.Post("postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require", headers, queryParams, map[string]string{"hello": "world"})
	if err == nil {
		t.Error("invalid url gave no error")
	}

	headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	_, err = client.Post("postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require", headers, queryParams, map[string][]string{"hello": {"world"}})
	if err == nil {
		t.Error("invalid url gave no error")
	}

	delete(headers, "Content-Type")
	_, err = client.Post("postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require", headers, queryParams, map[string]string{"hello": "world"})
	if err == nil {
		t.Error("invalid url gave no error")
	}
}

func TestClient_Post_NilData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/json"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Post(testServer.URL, headers, queryParams, nil)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}
}

func TestClient_Post_FormURLEncoded(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()
		response.Data = map[string]string{"hello": "world"}

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Post(testServer.URL, headers, queryParams, map[string][]string{"hello": {"world"}})
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Data.(map[string]interface{})["hello"] != "world" {
		t.Errorf("Expected: %s \n Got: %s", "world", testServerResp.Data.(map[string]string)["hello"])
	}
}

func TestClient_Post_FormURLEncoded_InvalidDataStructure(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()
		response.Data = map[string]string{"hello": "world"}

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	_, err := client.Post(testServer.URL, headers, queryParams, map[string]string{"hello": "world"})
	if err == nil {
		t.Error("Expects error for invalid data structure but did not get one")
		return
	}

}

func TestClient_Post_FormURLEncoded_NilData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Post(testServer.URL, headers, queryParams, nil)
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}
}

func TestClient_Put(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()
		response.Data = map[string]string{"hello": "world"}

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/json"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Put(testServer.URL, headers, queryParams, map[string]string{"hello": "world"})
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Data.(map[string]interface{})["hello"] != "world" {
		t.Errorf("Expected: %s \n Got: %s", "world", testServerResp.Data.(map[string]string)["hello"])
	}
}

func TestClient_Patch(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()
		response.Data = map[string]string{"hello": "world"}

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/json"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Patch(testServer.URL, headers, queryParams, map[string]string{"hello": "world"})
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Data.(map[string]interface{})["hello"] != "world" {
		t.Errorf("Expected: %s \n Got: %s", "world", testServerResp.Data.(map[string]string)["hello"])
	}
}

func TestClient_Delete(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := TestServerResponse{}

		response.Headers = request.Header
		response.Args = request.URL.Query()
		response.Data = map[string]string{"hello": "world"}

		responseBytes, _ := json.Marshal(response)

		writer.Write(responseBytes)
	}))

	headers := map[string][]string{
		"test": {"GET"},
		"Content-Type": {"application/json"},
	}

	queryParams := map[string][]string{
		"arg1": {"test"},
	}

	client := New()
	resp, err := client.Delete(testServer.URL, headers, queryParams, map[string]string{"hello": "world"})
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		testServer.Close()

		err := resp.Body.Close()
		if err != nil {
			t.Error(err)
		}
	}()

	var testServerResp TestServerResponse
	if err := json.NewDecoder(resp.Body).Decode(&testServerResp); err != nil {
		t.Error(err)
		return
	}

	if testServerResp.Args["arg1"][0] != "test" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Headers["Test"][0] != "GET" {
		t.Errorf("Expected: %s \n Got: %s", "test", testServerResp.Args["arg1"][0])
	}

	if testServerResp.Data.(map[string]interface{})["hello"] != "world" {
		t.Errorf("Expected: %s \n Got: %s", "world", testServerResp.Data.(map[string]string)["hello"])
	}
}
