package v4

import (
	"bytes"
	"net/http"
	"net/http/httputil"
)

// GetStringBodyHTTPRequest REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPRequest(r *http.Request) *string {
	if r == nil {
		return nil
	}

	headers, err := httputil.DumpRequest(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpRequest(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

	body := headersAndBody[len(headers):]
	s := string(bytes.TrimSpace(body))
	return &s
}

// GetStringBodyHTTPRequestJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPRequestJSON(r *http.Request) *string {
	if r == nil {
		return nil
	}

	headers, err := httputil.DumpRequest(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpRequest(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

	body := bytes.TrimSpace(headersAndBody[len(headers):])

	if len(body) > 0 {
		start := bytes.IndexAny(body, "{")
		end := bytes.LastIndexAny(body, "}")
		r := string(body[start : end+1])
		return &r
	}

	return nil
}

// GetStringBodyHTTPResponse REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPResponse(r *http.Response) *string {
	if r == nil {
		return nil
	}

	headers, err := httputil.DumpResponse(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpResponse(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

	body := headersAndBody[len(headers):]
	s := string(bytes.TrimSpace(body))
	return &s
}

// GetStringBodyHTTPResponseJSON REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetStringBodyHTTPResponseJSON(r *http.Response) *string {
	if r == nil {
		return nil
	}

	headers, err := httputil.DumpResponse(r, false)
	if err != nil {
		return nil
	}

	headersAndBody, err := httputil.DumpResponse(r, true)
	if err != nil || len(headersAndBody) == 0 {
		return nil
	}

	body := bytes.TrimSpace(headersAndBody[len(headers):])
	if len(body) > 0 {
		start := bytes.IndexAny(body, "{")
		end := bytes.LastIndexAny(body, "}")
		r := string(body[start : end+1])
		return &r
	}
	return nil
}
