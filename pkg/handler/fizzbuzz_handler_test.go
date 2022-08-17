package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestFizzbuzzHandler tests handlers of FizzBuzz web server
func TestFizzbuzzHandler(t *testing.T) {
	testSequence := []struct {
		request                  fizzbuzzRequest
		expectedPostResponse     []string
		expectedGetResponse      mostFrequentResponse
		expectedPostResponseCode int
	}{
		{
			request: fizzbuzzRequest{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedPostResponse: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
			expectedGetResponse: mostFrequentResponse{
				Count: 1,
				Requests: []fizzbuzzRequest{
					{
						Int1:  3,
						Int2:  5,
						Limit: 20,
						Str1:  "fizz",
						Str2:  "buzz",
					},
				},
			},
			expectedPostResponseCode: http.StatusOK,
		},
		{
			request: fizzbuzzRequest{
				Int1:  2,
				Int2:  4,
				Limit: 10,
				Str1:  "bad",
				Str2:  "good",
			},
			expectedPostResponse: []string{"1", "bad", "3", "badgood", "5", "bad", "7", "badgood", "9", "bad"},
			expectedGetResponse: mostFrequentResponse{
				Count: 1,
				Requests: []fizzbuzzRequest{
					{
						Int1:  3,
						Int2:  5,
						Limit: 20,
						Str1:  "fizz",
						Str2:  "buzz",
					},
					{
						Int1:  2,
						Int2:  4,
						Limit: 10,
						Str1:  "bad",
						Str2:  "good",
					},
				},
			},
			expectedPostResponseCode: http.StatusOK,
		},
		{
			request: fizzbuzzRequest{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedPostResponse: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
			expectedGetResponse: mostFrequentResponse{
				Count: 2,
				Requests: []fizzbuzzRequest{
					{
						Int1:  3,
						Int2:  5,
						Limit: 20,
						Str1:  "fizz",
						Str2:  "buzz",
					},
				},
			},
			expectedPostResponseCode: http.StatusOK,
		},
		{
			request: fizzbuzzRequest{
				Int1:  2,
				Int2:  4,
				Limit: 10,
				Str1:  "bad",
				Str2:  "good",
			},
			expectedPostResponse: []string{"1", "bad", "3", "badgood", "5", "bad", "7", "badgood", "9", "bad"},
			expectedGetResponse: mostFrequentResponse{
				Count: 2,
				Requests: []fizzbuzzRequest{
					{
						Int1:  3,
						Int2:  5,
						Limit: 20,
						Str1:  "fizz",
						Str2:  "buzz",
					},
					{
						Int1:  2,
						Int2:  4,
						Limit: 10,
						Str1:  "bad",
						Str2:  "good",
					},
				},
			},
			expectedPostResponseCode: http.StatusOK,
		},
		{
			request: fizzbuzzRequest{
				Int1:  2,
				Int2:  4,
				Limit: 10,
				Str1:  "bad",
				Str2:  "good",
			},
			expectedPostResponse: []string{"1", "bad", "3", "badgood", "5", "bad", "7", "badgood", "9", "bad"},
			expectedGetResponse: mostFrequentResponse{
				Count: 3,
				Requests: []fizzbuzzRequest{
					{
						Int1:  2,
						Int2:  4,
						Limit: 10,
						Str1:  "bad",
						Str2:  "good",
					},
				},
			},
			expectedPostResponseCode: http.StatusOK,
		},
	}

	fizzbuzzHandler := NewFizzbuzzHandler()
	router := gin.Default()
	router.POST("/fizzbuzz", fizzbuzzHandler.HandleFizzbuzz)
	router.GET("/most-frequent", fizzbuzzHandler.HandleMostFrequent)
	recorder := httptest.NewRecorder()

	for _, tc := range testSequence {
		// assert POST fizzbuzz request
		body, err := json.Marshal(tc.request)
		if err != nil {
			t.Error(err)
		}
		var responses []string
		err = checkRequest(router, recorder, "POST", "/fizzbuzz", body, tc.expectedPostResponseCode, &responses, &tc.expectedPostResponse)
		if err != nil {
			t.Error(err)
		}

		// assert GET frequent request
		var getResponse mostFrequentResponse
		err = checkRequest(router, recorder, "GET", "/most-frequent", nil, http.StatusOK, &getResponse, &tc.expectedGetResponse)
		if err != nil {
			t.Error(err)
		}
	}
}

// TestErrorHandling tests bad request scenario
func TestErrorHandling(t *testing.T) {
	fizzbuzzHandler := NewFizzbuzzHandler()
	router := gin.Default()
	router.POST("/fizzbuzz", fizzbuzzHandler.HandleFizzbuzz)
	recorder := httptest.NewRecorder()
	badRequestMsg := &errorMessage{Error: "missing arguments"}
	var responses errorMessage

	// check missing arguments
	err := checkRequest(router, recorder, "POST", "/fizzbuzz", []byte(`{"int1": 3, "int2": 5, "limit": 20, "str1": "fizz"}`), http.StatusBadRequest, &responses, badRequestMsg)
	if err != nil {
		t.Error(err)
	}

	// check zero integer
	err = checkRequest(router, recorder, "POST", "/fizzbuzz", []byte(`{"int1": 3, "int2": 0, "limit": 20, "str1": "fizz", "str2": "fuzz"}`), http.StatusBadRequest, &responses, badRequestMsg)
	if err != nil {
		t.Error(err)
	}
}

// checkRequest checks if the handling of a specific request matchs the expectation
func checkRequest(router *gin.Engine, recorder *httptest.ResponseRecorder, verb string, path string, requestBody []byte, expectedResponseCode int, response, expectedResponse any) error {
	req, err := http.NewRequest(verb, path, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	router.ServeHTTP(recorder, req)
	responseData, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(recorder.Code, expectedResponseCode) {
		return fmt.Errorf("Request returns %d code, while %d was expected", recorder.Code, expectedResponseCode)
	}

	fmt.Printf("%s", responseData)

	err = json.Unmarshal(responseData, response)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(response, expectedResponse) {
		return fmt.Errorf("Request returns %v, while %v was expected", response, expectedResponse)
	}

	return nil
}

// mostFrequentResponse encapsulates the response of the most-requent GET request
type mostFrequentResponse struct {
	Count    int               `json:"count"`
	Requests []fizzbuzzRequest `json:"requests"`
}

// errorMessage encapsulates the error message returned by handlers
type errorMessage struct {
	Error string `json:"error"`
}
