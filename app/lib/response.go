package api

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Code    int
	Message string
	Info    gin.H
}

type ErrorResponse struct {
	Code    int
	Message string
}
