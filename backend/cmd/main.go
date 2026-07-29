// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Hello from Go backend",
// 	})
// }

// func main() {
// 	http.HandleFunc("/api/hello", helloHandler)

// 	println("Server running on :8080")

// 	http.ListenAndServe(":8080", nil)
// }

// package main

// import "backend/internal/database"

// func main() {

// 	database.ConnectDB()

// }

package main

import (
	"log"
	"net/http"
	
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"backend/internal/database"
	"backend/internal/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			"http://localhost:5173",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
