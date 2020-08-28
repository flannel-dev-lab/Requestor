package requestor

import (
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
