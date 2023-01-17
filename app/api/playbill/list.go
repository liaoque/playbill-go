package playbill

import (
	"github.com/gin-gonic/gin"
	"palybill/app/model"
)

func listHandler(c *gin.Context) {
	poster := model.Poster{}
	getAll := poster.GetAll(c)
	return success
}
