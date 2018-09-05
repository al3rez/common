package test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func DoRequest(server *server, method, path string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	res := httptest.NewRecorder()
	server.mux.ServeHTTP(res, req)
	return res, nil
}
