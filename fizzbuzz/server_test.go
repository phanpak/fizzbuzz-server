package fizzbuzz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostFizzBuzz(t *testing.T) {
	server := NewServer()
	server.registerRoutes()

	// Create a test server
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	// Create a request body
	body := `{"int1": 3, "int2": 5, "limit": 15, "str1": "fizz", "str2": "buzz"}`

	// Send a POST request to the test server
	resp, err := http.Post(ts.URL+"/fizzbuzz", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// assert that status is OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// Read the response body
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the response body
	respBody := struct {
		Result []string `json:"result"`
	}{}
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		// t.Fatal(err)
	}

	expectedResult := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	if !sliceEqual(respBody.Result, expectedResult) {
		t.Errorf("response body: got %v, want %v", respBody.Result, expectedResult)
	}
}

func TestGetStats(t *testing.T) {
	server := NewServer()
	server.registerRoutes()

	// Create a test server
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	// Populate the hit counter
	body := `{"int1": 3, "int2": 5, "limit": 15, "str1": "fizz", "str2": "buzz"}`

	// Send a POST request to the test server
	resp, err := http.Post(ts.URL+"/fizzbuzz", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// assert that status is OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// Send a GET request to the test server
	resp, err = http.Get(ts.URL + "/stats")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// assert that status is OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// Read the response body
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("strBody:", string(respBodyBytes))

	// Unmarshal the response body
	respBody := struct {
		Hits int `json:"hits"`
	}{}
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the response
	if respBody.Hits != 1 {
		t.Errorf("response body: got %v, want %v", respBody.Hits, 1)
	}
}

func TestGetStats_Empty(t *testing.T) {
	server := NewServer()
	server.registerRoutes()

	// Create a test server
	ts := httptest.NewServer(server.router)
	defer ts.Close()

	// Send a GET request to the test server
	resp, err := http.Get(ts.URL + "/stats")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// assert that status is OK
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// Read the response body
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal the response body
	respBody := struct {
		Total int `json:"total"`
	}{}
	err = json.Unmarshal(respBodyBytes, &respBody)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the response
	if respBody.Total != 0 {
		t.Errorf("response body: got %v, want %v", respBody.Total, 1)
	}
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
