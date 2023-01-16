package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context) {
	//	return AppResponse\AppResponse::success([
	//		'token' => 'eyJhbGciOiJIUzUxMiJ9.test'
	//]);
	c.JSON(http.StatusOK, gin.H{
		"token": "eyJhbGciOiJIUzUxMiJ9.test",
	})
	//c.String()
}
