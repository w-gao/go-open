package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// SmartOpen smartly opens a file based on the URI. Supported URIs are file://,
// http://, and https://.
//
// Read() of the returned io.ReadCloser may raise an error. For example, when
// the connection terminates when streaming from the Internet.
//
// It is also the caller's responsibility to close the returned io.ReadCloser.
func SmartOpen(uri string) (io.ReadCloser, error) {

	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		// TODO: options (headers, etc) from user
		resp, err := http.DefaultClient.Get(uri)

		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return resp.Body, nil
	}

	if strings.HasPrefix(uri, "file://") {
		uri = uri[7:]
	}

	// Everything else falls back to os.Open()
	return os.Open(uri) // f, err
}
