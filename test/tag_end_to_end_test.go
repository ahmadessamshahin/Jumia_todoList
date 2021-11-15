package test

import (
	"Jumia_todoList/api/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateTag(t *testing.T) {
	gin.SetMode(gin.TestMode)
	data := url.Values{}
	data.Set("name", "todo")
	// data.Set("task_id", task.ID)
	req := httptest.NewRequest(http.MethodPost, "/tag", strings.NewReader(data.Encode()))
	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
	var m model.TagCreateOutput
	err := json.Unmarshal(response.Body.Bytes(), &m)
	if err != nil {
		t.Errorf("Miss Matching Response %s", response.Body.String())
	}
}
