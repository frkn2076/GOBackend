package test

// import (
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pact-foundation/pact-go/dsl"
// )

// func startServer() {
// 	router := gin.Default()

// 	todoRoute := router.Group("/todo")
// 	{
// 		todoRoute.POST("add", todoController.Add)
// 		todoRoute.GET("items", todoController.Items)
// 	}

// 	router.Run(":5000")
// }

// func TestProvider(t *testing.T) {

// 	// Create Pact connecting to local Daemon
// 	pact := &dsl.Pact{
// 		Provider: "MyProvider",
// 	}

// 	// Start provider API in the background
// 	go startServer()

// 	// Verify the Provider using the locally saved Pact Files
// 	pact.VerifyProvider(t, dsl.types.VerifyRequest, {
// 		ProviderBaseURL: "http://localhost:5000",
// 		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("%s/myconsumer-myprovider.json", pactDir))},
// 		StateHandlers: types.StateHandlers{
// 			// Setup any state required by the test
// 			// in this case, we ensure there is a "user" in the system
// 			"User foo exists": func() error {
// 				lastName = "crickets",
// 				return nil
// 			},
// 		},
// 	})
// }
