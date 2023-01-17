package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	l "palybill/app/lib"
)

func loginHandler(c *gin.Context) {
	//	return AppResponse\AppResponse::success([
	//		'token' => 'eyJhbGciOiJIUzUxMiJ9.test'
	//]);
	c.JSON(http.StatusOK, l.SuccessResponse{
		Code:    0,
		Message: "",
		Info: gin.H{
			"token": "eyJhbGciOiJIUzUxMiJ9.test",
		},
	})
	//c.String()
}
