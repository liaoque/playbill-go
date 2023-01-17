package api

import "github.com/gin-gonic/gin"

type successResponse struct {
	code    int
	message string
	info    gin.H
}

type errorResponse struct {
	code    int
	message string
}
