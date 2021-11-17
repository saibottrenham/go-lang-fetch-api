package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

// unit test for function sender
func TestSender(t *testing.T) {
	// mock the http request that get's sent subsequntialy
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("POST", "http://localhost:8082/receiver",
		httpmock.NewStringResponder(200, "OK"))

	// test payload example
	test_payload := `{
		"to" : [
				{"email": "tobiasmahnert@gmail.com"},
				{"email": "tobiasmahnert@web.de"}
		],
		"from" : {
			"email_address_id": 16426,
			"name": "tobias"
		},
		"schedule": 1624166184,
		"subject": "test subject",
		"body": "test body"
	}`
	// test sender request function
	req := httptest.NewRequest(http.MethodGet, "/sender", strings.NewReader(test_payload))
	w := httptest.NewRecorder()
	sender(w, req)
	res := w.Result()
	defer res.Body.Close()
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if httpmock.GetTotalCallCount() != 1 {
		t.Errorf("expected call count to be 1 got %v", httpmock.GetTotalCallCount())
	}
}

func TestReceiver(t *testing.T) {
	httpmock.DeactivateAndReset()
	// test payload example
	test_payload := `{
		"to" : [
				{"email": "tobiasmahnert@gmail.com"},
				{"email": "tobiasmahnert@web.de"}
		],
		"from" : {
			"email_address_id": 16426,
			"name": "tobias"
		},
		"schedule": 1624166184,
		"subject": "test subject",
		"body": "test body"
	}`
	// test sender request function
	req := httptest.NewRequest(http.MethodPost, "/receiver", strings.NewReader(test_payload))
	w := httptest.NewRecorder()
	receiver(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "tobias" {
		t.Errorf("expected status to be tobias got %v", string(data))
	}
}
