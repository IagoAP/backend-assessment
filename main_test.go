package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	payload := []byte(`{"Username":"User1","Password":"User1"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
	}
}

func TestNotLogin(t *testing.T) {
	payload := []byte(`{"Username":"Wrong","Password":"Wrong"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected %d, received %d", http.StatusUnauthorized, res.StatusCode)
	}
}

func TestActivationRequests(t *testing.T) {
	payload := []byte(`{"Username":"User1","Password":"User1"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var token string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/ActivationRequests", nil)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func TestIssueProductActivation(t *testing.T) {
	payload := []byte(`{"Username":"Empresa1","Password":"Empresa1"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var token string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	payload = []byte(`{"Description":"Teste","CustomerMid":2,"CustomerEmail":"iagoaph@gmail.com"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/IssueProductActivation", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func TestNotIssueProductActivation(t *testing.T) {
	payload := []byte(`{"Username":"User1","Password":"User1"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var token string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	payload = []byte(`{"Description":"Teste","CustomerMid":2,"CustomerEmail":"iagoaph@gmail.com"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/IssueProductActivation", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected %d, received %d", http.StatusUnauthorized, res.StatusCode)
	}
}

func TestApproveActivation(t *testing.T){
	testApproveRejectActivation(t, "http://localhost:8080/ApproveActivation")
}

func TestRejectActivation(t *testing.T){
	testApproveRejectActivation(t, "http://localhost:8080/RejectActivation")
}

func testApproveRejectActivation(t *testing.T, url string) {
	payload := []byte(`{"Username":"Empresa1","Password":"Empresa1"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var token string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	payload = []byte(`{"Description":"Teste","CustomerMid":2,"CustomerEmail":"iagoaph@gmail.com"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/IssueProductActivation", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var ActivationID string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["ActivationID"] == "" {
			t.Errorf("Empty ID")
		}
		ActivationID = m["ActivationID"]
	}

	payload = []byte(`{"Username":"User1","Password":"User1"}`)
	res, err = http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	payload = []byte(`{"ID":"` + ActivationID + `"}`)

	client = &http.Client{}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}

func TestNotApproveActivation(t *testing.T){
	testNotApproveRejectActivation(t, "http://localhost:8080/ApproveActivation")
}

func TestNotRejectActivation(t *testing.T){
	testNotApproveRejectActivation(t, "http://localhost:8080/RejectActivation")
}

func testNotApproveRejectActivation(t *testing.T, url string) {
	payload := []byte(`{"Username":"Empresa1","Password":"Empresa1"}`)
	res, err := http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var token string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	payload = []byte(`{"Description":"Teste","CustomerMid":2,"CustomerEmail":"iagoaph@gmail.com"}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/IssueProductActivation", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	var ActivationID string

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["ActivationID"] == "" {
			t.Errorf("Empty ID")
		}
		ActivationID = m["ActivationID"]
	}

	payload = []byte(`{"Username":"Empresa1","Password":"Empresa1"}`)
	res, err = http.Post("http://localhost:8080/RequestToken", "application/json", bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		var m map[string]string
		err = json.Unmarshal(bodyBytes, &m)
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if m["access_token"] == "" {
			t.Errorf("Empty token")
		}
		token = m["access_token"]
	}

	payload = []byte(`{"ID":"` + ActivationID + `"}`)

	client = &http.Client{}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	req.Header.Set("Token", token)
	res, err = client.Do(req)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Errorf("Expected %d, received %d", http.StatusUnauthorized, res.StatusCode)
	}
}
