package list

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	listMocks "Jumia_todoList/mocks/usecase/list"
	"Jumia_todoList/test/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
)

func TestListCreation(t *testing.T) {

	r, listMockUseCase := setupListRouter()

	requestBuildFunc := func(body map[string]interface{}) (*http.Request, error) {
		jsonData, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		return http.NewRequest(http.MethodPost, "/list/", bytes.NewReader(jsonData))
	}
	mockingFunc := func() {
		l := model.ListCreateInput{Name: "Test"}
		listMockUseCase.On("Create", l).Return(1, nil).Once()
	}

	cases := []helper.CaseHandler{
		{
			Name: "ValidListCreation",
			Params: map[string]interface{}{
				"name": "Test",
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusCreated, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"message":"success","data":{"id":1}}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
		{
			Name: "InValidParamtersListCreation",
			Params: map[string]interface{}{
				"name": 1,
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusBadRequest, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"error":"INVALID INPUT"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
	}

	helper.CaseHandlerEntryIterator(t, cases, r)

}

func TestListCreationServiceFailure(t *testing.T) {

	r, listMockUseCase := setupListRouter()

	requestBuildFunc := func(body map[string]interface{}) (*http.Request, error) {
		jsonData, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		return http.NewRequest(http.MethodPost, "/list/", bytes.NewReader(jsonData))
	}

	errorMockingFunc := func() {
		l := model.ListCreateInput{Name: "Test"}
		listMockUseCase.On("Create", l).Return(0, fmt.Errorf("LIST USECASE FAILURE")).Once()
	}

	cases := []helper.CaseHandler{
		{
			Name: "ListCreationUseCaseFailure",
			Params: map[string]interface{}{
				"name": "Test",
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      errorMockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusBadRequest, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"error":"LIST USECASE FAILURE"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
	}
	helper.CaseHandlerEntryIterator(t, cases, r)
}

func TestListUpdate(t *testing.T) {

	r, listMockUseCase := setupListRouter()

	requestBuildFunc := func(body map[string]interface{}) (*http.Request, error) {
		jsonData, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		return http.NewRequest(http.MethodPatch, "/list/", bytes.NewReader(jsonData))
	}

	mockingFunc := func() {
		l := model.ListUpdateInput{ID: 1, Name: "Test1"}
		listMockUseCase.On("Update", l).Return(nil).Once()
	}

	cases := []helper.CaseHandler{
		{
			Name: "ValidListUpdate",
			Params: map[string]interface{}{
				"id":   1,
				"name": "Test1",
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusOK, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"message":"success"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
		{
			Name: "InValidParamtersListUpdate",
			Params: map[string]interface{}{
				"id":   "tt",
				"name": 3,
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusBadRequest, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"error":"INVALID INPUT"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
	}

	helper.CaseHandlerEntryIterator(t, cases, r)
}

func TestListDelete(t *testing.T) {

	r, listMockUseCase := setupListRouter()

	requestBuildFunc := func(body map[string]interface{}) (*http.Request, error) {
		jsonData, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		return http.NewRequest(http.MethodDelete, "/list/", bytes.NewReader(jsonData))
	}

	mockingFunc := func() {
		l := model.ListRemoveInput{ID: 1}
		listMockUseCase.On("Delete", l).Return(nil).Once()
	}

	cases := []helper.CaseHandler{
		{
			Name: "ValidListDelete",
			Params: map[string]interface{}{
				"id": 1,
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusOK, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"message":"success"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
		{
			Name: "InValidParamtersListDelete",
			Params: map[string]interface{}{
				"id": "tt",
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusBadRequest, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"error":"INVALID INPUT"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
	}

	helper.CaseHandlerEntryIterator(t, cases, r)
}

func TestListDeleteServiceFailure(t *testing.T) {

	r, listMockUseCase := setupListRouter()

	requestBuildFunc := func(body map[string]interface{}) (*http.Request, error) {
		jsonData, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		return http.NewRequest(http.MethodDelete, "/list/", bytes.NewReader(jsonData))
	}
	errorMockingFunc := func() {
		l := model.ListRemoveInput{ID: 1}
		listMockUseCase.On("Delete", l).Return(fmt.Errorf("LIST USECASE FAILURE")).Once()
	}
	cases := []helper.CaseHandler{
		{
			Name: "ListDeleteUseCaseFailure",
			Params: map[string]interface{}{
				"id": 1,
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      errorMockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusBadRequest, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"error":"LIST USECASE FAILURE"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
	}
	helper.CaseHandlerEntryIterator(t, cases, r)
}

func TestListGets(t *testing.T) {

	r, listMockUseCase := setupListRouter()

	requestBuildFunc := func(queryParams map[string]interface{}) (*http.Request, error) {
		return http.NewRequest(http.MethodGet, fmt.Sprintf("/list/all?limit=%v&offset=%v", queryParams["limit"], queryParams["offset"]), nil)
	}

	mockingFunc := func() {
		l := model.GetAllListInput{Limit: 10, Offset: 0}
		listMockUseCase.On("GetAllLists", l).Return([]entity.List{{ID: 1, Name: "Test"}}).Once()
	}

	cases := []helper.CaseHandler{
		{
			Name: "ValidListGet",
			Params: map[string]interface{}{
				"limit":  10,
				"offset": 0,
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusOK, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"message":"success","data":[{"id":1,"name":"Test"}]}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
		{
			Name: "InValidParamtersListGet",
			Params: map[string]interface{}{
				"limit":  "a",
				"offset": 0,
			},
			RequestBuildFunc: requestBuildFunc,
			MockingFunc:      mockingFunc,
			ValidationTesting: []func(w *httptest.ResponseRecorder) bool{
				func(w *httptest.ResponseRecorder) bool {
					return assert.Equal(t, http.StatusBadRequest, w.Code)
				},
				func(w *httptest.ResponseRecorder) bool {
					expected := `{"error":"INVALID INPUT"}`
					return assert.Equal(t, expected, w.Body.String())
				},
			},
		},
	}

	helper.CaseHandlerEntryIterator(t, cases, r)
}

func setupListRouter() (*gin.Engine, *listMocks.UseCase) {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	usecase := new(listMocks.UseCase)
	handlerRouter := &GinHandler{}
	handlerRouter.Routes(r, usecase)

	return r, usecase
}
