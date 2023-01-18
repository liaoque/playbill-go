package playbill

import (
	"github.com/gin-gonic/gin"
	"net/http"
	l "palybill/app/lib"
	"palybill/app/model"
)

func listHandler(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	poster := model.Poster{}
	all := poster.GetAll(c, keyword, 1, 10)
	c.JSON(http.StatusOK, l.SuccessResponse{
		Code:    0,
		Message: "",
		Info:    all,
	})
}
