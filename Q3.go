package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type CustomRequest struct {
	BoolField   bool    `json:"bool_field"`
	IntField    int     `json:"int_field"`
	StringField string  `json:"string_field"`
	RuneField   rune    `json:"rune_field"`
	PointerField *string `json:"pointer_field"`
}

func main() {
	router := gin.Default()

	router.POST("/custom", func(c *gin.Context) {
		var req CustomRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("Received Custom Request: %+v\n", req)
		c.JSON(200, gin.H{"message": "Request processed successfully"})
	})

	router.Run(":8000")
}
