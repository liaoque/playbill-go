package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Global() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Writer.Status()
		if status == 200 {
			fmt.Println(status)
			//c.Writer.WriteString()
			c.JSON(http.StatusOK, gin.H{"sdsada": "1111"})

		}

		c.Next()

	}
}
