package test

import (
	"Jumia_todoList/api/handler"
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var router = handler.Handler.Router

var list = entity.List{}

func TestCreateList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	data := url.Values{}
	data.Set("name", "foo")
	req := httptest.NewRequest(http.MethodPost, "/list", strings.NewReader(data.Encode()))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m model.ListCreateOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
	list.ID = int64(m.Data.ID)
}

func TestUpdateList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	data := url.Values{}
	data.Set("name", "Boo")
	// data.Set("id", list.ID)
	req := httptest.NewRequest(http.MethodPatch, "/list", strings.NewReader(data.Encode()))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m model.EmptySuccessfulOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
}

func TestGetAllLists(t *testing.T) {
	gin.SetMode(gin.TestMode)
	req := httptest.NewRequest(http.MethodGet, "/list/all", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m model.GetAllListOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}

	if m.Data == nil || len(m.Data) == 0 {
		t.Errorf("Miss Matching Response %v", m.Data)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
