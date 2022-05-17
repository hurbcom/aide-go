package aidego

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

// GetByteArrayAndBufferFromRequestBody
func GetByteArrayAndBufferFromRequestBody(body io.ReadCloser) ([]byte, *bytes.Buffer, error) {
	defer body.Close()
	byteArray, err := ioutil.ReadAll(body)
	if err != nil {
		return []byte{}, nil, err
	}
	buffer := bytes.NewBuffer(byteArray)
	return byteArray, buffer, nil
}

func ThisBytesContains(b []byte, s string) bool {
	return strings.Contains(strings.ToLower(string(b)), s)
}
