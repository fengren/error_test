package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Panic Demo",
			"routes": gin.H{
				"/panic-array":     "Trigger array index out of range panic",
				"/panic-nil":       "Trigger nil pointer dereference panic",
				"/panic-divide":    "Trigger divide by zero panic",
				"/panic-map":       "Trigger map key not found panic",
				"/panic-interface": "Trigger interface conversion panic",
			},
		})
	})

	r.GET("/panic-array", func(c *gin.Context) {
		arr := []int{1, 2, 3}
		value := arr[10]
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.GET("/panic-nil", func(c *gin.Context) {
		var ptr *int
		value := *ptr
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.GET("/panic-map", func(c *gin.Context) {
		m := make(map[string]int)
		value := m["nonexistent"]
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.GET("/panic-interface", func(c *gin.Context) {
		var i interface{} = "string"
		_ = i.(int)
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	r.Run(":8080")
}
