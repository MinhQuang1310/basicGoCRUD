package main_test

// import (
// 	"basicGoCrud/db"
// 	"basicGoCrud/handlers"
// 	"net/http"
// 	"testing"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestMain(m *testing.M) {
// 	// Run the tests
// 	m.Run()

// 	// Clean up the database connection
// 	if err := db.CleanDB(db.InitDB()); err != nil {
// 		panic(err)
// 	}
// }

// func TestMainFunction(t *testing.T) {
// 	// Create a new router instance
// 	router := gin.Default()

// 	// Set up the routes
// 	v1 := router.Group("/v1")
// 	{
// 		items := v1.Group("/items")
// 		{
// 			// Handle POST requests to create a new todo item
// 			items.POST("", handlers.CreateItemHandler(db.InitDB()))
// 		}
// 	}

// 	// Start the server on port 8080
// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: router,
// 	}

// 	// Channel to communicate errors from the goroutine
// 	errChan := make(chan error)

// 	// Run the server in a separate goroutine
// 	go func() {
// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			errChan <- err
// 		}
// 		close(errChan)
// 	}()

// 	// Give the server a moment to start
// 	time.Sleep(100 * time.Millisecond)

// 	// Make a POST request to create a new todo item
// 	resp, err := http.Post("http://localhost:8080/v1/items", "application/json", nil)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusCreated, resp.StatusCode)

// 	// Stop the server
// 	server.Shutdown(nil)

// 	// Check for errors from the server goroutine
// 	if err := <-errChan; err != nil {
// 		t.Fatal(err)
// 	}
// }
