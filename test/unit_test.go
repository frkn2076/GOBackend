package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/GOBackend/controllers"
	"app/GOBackend/controllers/request"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddItemSuccessfully(t *testing.T) {
	// Arrange
	responseWriter := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseWriter)

	request := request.TodoRequest{Name: "Reading book"}
	requestBytes, _ := json.Marshal(request)
	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(requestBytes))

	// Act
	controllers.AddItem(context)

	// Assert
	var responseBody map[string]interface{}
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, responseWriter.Code, 200)
	assert.Equal(t, responseBody["IsError"], false)
	assert.Nil(t, responseBody["ErrorMessage"])
}

func TestAddItemFails(t *testing.T) {
	// Arrange
	responseWriter := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseWriter)

	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer([]byte("")))

	// Act
	controllers.AddItem(context)

	// Assert
	var responseBody map[string]interface{}
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, responseWriter.Code, 200)
	assert.Equal(t, responseBody["IsError"], true)
	assert.Equal(t, responseBody["ErrorMessage"], "Geçici süre hizmet veremiyoruz")
}

func TestItemsSuccessfully(t *testing.T) {
	// Arrange
	responseWriter := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseWriter)

	context.Request, _ = http.NewRequest(http.MethodGet, "/items", bytes.NewBuffer([]byte("")))

	// Act
	controllers.Items(context)

	// Assert
	var responseBody map[string]interface{}
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	assert.Equal(t, responseWriter.Code, 200)
	assert.Equal(t, responseBody["IsError"], false)
	assert.NotNil(t, responseBody["Items"])
}
