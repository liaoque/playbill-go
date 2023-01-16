package api

import (
	"github.com/gin-gonic/gin"
	"palybill/app/middleware"
)

func Routers(r *gin.Engine) {
	v1 := r.Group("api")
	v1.Use(middleware.Global())
	v1.GET("/login", loginHandler)

	//v1.GET("/getAsyncRoutes", loginHandler)
	//v1.GET("/uploads", loginHandler)
	//v1.GET("/playbill/view", loginHandler)
	//v1.GET("/playbill/list", loginHandler)
	//v1.GET("/playbill/save", loginHandler)
	//v1.GET("/playbill/query/id/:id", loginHandler)
	//v1.GET("/error/403.html", loginHandler)
	//v1.GET("/error/502.html", loginHandler)

	//// 2.绑定路由规则，执行的函数
	//// gin.Context，封装了request和response
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/login", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/getAsyncRoutes", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/uploads", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/playbill/view", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/playbill/list", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/playbill/save", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/playbill/query/id/:id", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/error/403.html", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//
	//r.GET("/api/error/502.html", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
}
