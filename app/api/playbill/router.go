package playbill

import "github.com/gin-gonic/gin"

func Routers(r *gin.Engine) {
	v1 := r.Group("api/playbill")
	v1.GET("view", viewHandler)
	v1.GET("list", listHandler)
	v1.GET("save", saveHandler)
	v1.GET("query/id/:id", queryHandler)
}
