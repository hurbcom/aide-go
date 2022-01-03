package aidego

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetStringBodyHTTPRequest(t *testing.T) {
	body, _ := json.Marshal(nil)
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(body))
	actual := GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "null", *actual)

	req, _ = http.NewRequest("POST", "http://server.com", nil)
	actual = GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "", *actual)

	req, _ = http.NewRequest("POST", "http://server.com", nil)
	req.Header = nil
	actual = GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "", *actual)

	req, _ = http.NewRequest("POST", "http://server.com", bytes.NewBuffer([]byte("PLAIN TEXT")))
	req.Header = nil
	actual = GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "PLAIN TEXT", *actual)
}

func TestGetStringBodyHTTPRequestJSON(t *testing.T) {
	body, _ := json.Marshal(map[string]string{"foo": "bar"})
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(body))
	actual := GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "{\"foo\":\"bar\"}", *actual)
}

func TestGetStringBodyHTTPRequestPlainText(t *testing.T) {
	stringBody := "PLAIN TEXT"
	byteArrayStringBody := []byte(stringBody)
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(byteArrayStringBody))
	actual := GetStringBodyHTTPRequest(req)

	assert.NotNil(t, actual)
	assert.Equal(t, stringBody, *actual)
}

func TestGetStringBodyHTTPRequestJSONEncoded(t *testing.T) {
	stringBody := `1223ab
{'response':{'code':200}}
0

`
	byteArrayStringBody := []byte(stringBody)
	req, _ := http.NewRequest("POST", "http://server.com", bytes.NewBuffer(byteArrayStringBody))
	actual := GetStringBodyHTTPRequestJSON(req)

	assert.NotNil(t, actual)
	assert.Equal(t, "{'response':{'code':200}}", *actual)
}

func TestGetStringBodyHTTPResponseJSON(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	defer gock.Clean()

	gock.New("http://server.com").
		Get("/bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	req, _ := http.NewRequest("GET", "http://server.com/bar", nil)
	client := &http.Client{}
	res, _ := client.Do(req)
	actual := GetStringBodyHTTPResponse(res)

	assert.NotNil(t, actual)
	assert.Equal(t, "{\"foo\":\"bar\"}", *actual)
}

func TestGetStringBodyHTTPResponsePlainText(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	defer gock.Clean()

	stringBody := "PLAIN TEXT"

	gock.New("http://server.com").
		Get("/bar").
		Reply(200).
		BodyString(stringBody)

	req, err := http.NewRequest("GET", "http://server.com/bar", nil)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	actual := GetStringBodyHTTPResponse(res)

	assert.NotNil(t, actual)
	assert.Equal(t, stringBody, *actual)
}

func TestGetStringBodyHTTPResponseJSONEncoded(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	defer gock.Clean()

	stringBody := `1223ab
{'response':{'code':200}}
0

`

	gock.New("http://server.com").
		Get("/bar").
		Reply(200).
		BodyString(stringBody)

	req, err := http.NewRequest("GET", "http://server.com/bar", nil)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nErr: %+v\n\n", err)
	}
	actual := GetStringBodyHTTPResponseJSON(res)

	assert.NotNil(t, actual)
	assert.Equal(t, "{'response':{'code':200}}", *actual)
}
