package test

import (
	"Jumia_todoList/api/model"
	"Jumia_todoList/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var task = entity.Task{}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var body = model.TaskCreateInput{
		ListID:      uint(list.ID),
		Title:       "finalize",
		Description: "",
		Due:         "2021-11-01 20:12:18.13069 +0200 EET m=+0.000128166",
		Completed:   false,
	}

	jsonBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(jsonBytes))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m model.TaskCreateOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
	task.ID = int64(m.Data.ID)
}

func TestUpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var body = model.TaskUpdateInput{
		ListID:      uint(list.ID),
		Title:       "finalize",
		Description: "",
		Due:         "2021-11-01 20:12:18.13069 +0200 EET m=+0.000128166",
		Completed:   false,
	}

	jsonBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPatch, "/task", bytes.NewReader(jsonBytes))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, response.Code)
	var m model.EmptySuccessfulOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
}

func TestGetListTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/task?id=%d", list.ID), nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, response.Code)
	var m model.GetListTask
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
	if m.Data == nil || len(m.Data) == 0 {
		t.Errorf("Miss Matching Response %v", m.Data)
	}
}

func TestFilterListTasksByTags(t *testing.T) {
	gin.SetMode(gin.TestMode)

	var body = model.TaskFilterInput{
		ListId: int(list.ID),
		Tags:   []string{"todo"},
	}

	jsonBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodGet, "/task/filter", bytes.NewReader(jsonBytes))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, response.Code)
	var m model.TaskFilterOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
	if m.Data == nil || len(m.Data) == 0 {
		t.Errorf("Miss Matching Response %v", m.Data)
	}
}
