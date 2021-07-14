package test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"

	"app/GOBackend/controllers"
	"app/GOBackend/database"
	"app/GOBackend/controllers/request"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddItem(t *testing.T) {
	// Arrange
	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	router.POST("add", controllers.AddItem)

	request := request.TodoRequest{Name: "Reading book"}
	requestBytes, _ := json.Marshal(request)

	context.Request, _ = http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(requestBytes))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	var responseBody map[string]interface{}
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, responseWriter.Code, 200)
	assert.Equal(t, responseBody["IsError"], false)
	assert.Nil(t, responseBody["ErrorMessage"])

	//delete last test insert
	database.RemoveLastTodo()
}

func TestItems(t *testing.T) {
	// Arrange
	responseWriter := httptest.NewRecorder()
	context, router := gin.CreateTestContext(responseWriter)

	router.GET("items", controllers.Items)

	context.Request, _ = http.NewRequest(http.MethodGet, "/items", bytes.NewBuffer([]byte("")))

	// Act
	router.ServeHTTP(responseWriter, context.Request)

	var responseBody map[string]interface{}
	if err := json.Unmarshal(responseWriter.Body.Bytes(), &responseBody); err != nil {
		panic(err)
	}

	// Assert
	assert.Equal(t, responseWriter.Code, 200)
	assert.Equal(t, responseBody["IsError"], false)
	assert.NotNil(t, responseBody["Items"])
}
