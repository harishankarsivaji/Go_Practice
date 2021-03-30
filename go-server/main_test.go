package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_GetTime(t *testing.T) {

	//making a http request
	req := httptest.NewRequest(http.MethodGet, "/time", nil)

	//make respone recorder for get response
	rw := httptest.NewRecorder()

	//expected result
	expectedResult := time.Now()

	//callimg the api
	GetTime(rw, req)

	response := rw.Result()

	body, _ := ioutil.ReadAll(response.Body)

	// because body is []byte, we have to change to string to compare with expectedResult
	if expectedResult != string(body) {
		t.Errorf("should be %s but is %s", expectedResult, body)
	}
}
