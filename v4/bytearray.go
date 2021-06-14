package v4

import (
	"bytes"
	"io"
	"io/ioutil"
)

// GetByteArrayAndBufferFromRequestBody REQUIRE THEM TO DOCUMENT THIS FUNCTION
func GetByteArrayAndBufferFromRequestBody(body io.ReadCloser) ([]byte, *bytes.Buffer, error) {
	defer body.Close()
	byteArray, err := ioutil.ReadAll(body)
	if err != nil {
		return []byte{}, nil, err
	}
	buffer := bytes.NewBuffer(byteArray)
	return byteArray, buffer, nil
}
