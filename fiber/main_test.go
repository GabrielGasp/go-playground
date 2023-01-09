package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  map[string]any
	}{
		{
			description:   "index route",
			route:         "/",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  map[string]any{"message": "OK"},
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  map[string]any{"error": "Not found"},
		},
	}

	// Setup the app as it is done in the main function
	app := fiberSetup()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req := httptest.NewRequest(
			http.MethodGet,
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, _ := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, false, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := io.ReadAll(res.Body)

		var bodyMap map[string]any
		json.Unmarshal(body, &bodyMap)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, bodyMap, test.description)
	}
}
