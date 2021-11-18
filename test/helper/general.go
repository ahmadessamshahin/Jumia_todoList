package helper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type CaseHandler struct {
	Name              string
	Params            map[string]interface{}
	RequestBuildFunc  func(map[string]interface{}) (*http.Request, error)
	MockingFunc       func()
	ValidationTesting []func(w *httptest.ResponseRecorder) bool
}

func CaseHandlerEntryIterator(t *testing.T, cases []CaseHandler, r *gin.Engine) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			c.MockingFunc()
			if c.RequestBuildFunc == nil {
				t.Fatal("MISSING RequestBuildFunc")
			}
			req, _ := c.RequestBuildFunc(c.Params)
			TestHTTPResponse(t, r, req, c.ValidationTesting...)
		})
	}
}

func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, funcs ...func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)
	for _, f := range funcs {
		if !f(w) {
			t.Fail()
		}
	}
}

type CaseService struct {
	Name              string
	Params            map[string]interface{}
	ServiceBuildFunc  func(map[string]interface{}) []interface{}
	MockingFunc       func()
	ValidationTesting []func(...interface{}) bool
}

func CaseServiceEntryIterator(t *testing.T, cases []CaseService) {
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			c.MockingFunc()
			if c.ServiceBuildFunc == nil {
				t.Fatal("MISSING RequestBuildFunc")
			}
			res := c.ServiceBuildFunc(c.Params)
			TestServiceResponse(t, res, c.ValidationTesting...)
		})
	}
}

func TestServiceResponse(t *testing.T, res []interface{}, funcs ...func(...interface{}) bool) {

	for _, f := range funcs {
		if !f(res...) {
			t.Fail()
		}
	}
}
