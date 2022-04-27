package tests

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/WeDias/golang-test-api/server"
	"github.com/stretchr/testify/assert"
)

var app = server.New()

func TestRoutes(t *testing.T) {
	tests := []struct {
		testName     string
		route        string
		method       string
		body         io.Reader
		expectedCode int
		expectedBody string
	}{
		{
			testName:     "an invalid path",
			route:        "/api/v1/invalid-path",
			method:       "GET",
			body:         nil,
			expectedCode: 404,
			expectedBody: "{\"error\":\"Cannot GET /api/v1/invalid-path\"}",
		},
		{
			testName:     "another invalid path",
			route:        "/api/v1/another-invalid-path",
			method:       "POST",
			body:         nil,
			expectedCode: 404,
			expectedBody: "{\"error\":\"Cannot POST /api/v1/another-invalid-path\"}",
		},
		{
			testName:     "a valid path",
			route:        "/api/v1/product",
			method:       "GET",
			body:         nil,
			expectedBody: "",
			expectedCode: 200,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			test.body,
		)

		res, err := app.Test(req, -1)
		assert.Nilf(t, err, test.testName)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.testName)

		if test.expectedBody != "" {
			body, err := ioutil.ReadAll(res.Body)
			assert.Nilf(t, err, test.testName)
			assert.Equalf(t, test.expectedBody, string(body), test.testName)
		}

	}
}
