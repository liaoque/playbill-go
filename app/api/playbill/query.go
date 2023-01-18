package playbill

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"palybill/app/lib"
	"palybill/app/model"
)

func queryHandler(c *gin.Context) {
	id := c.Param("id") //
	if id == "" {
		c.JSON(http.StatusOK, lib.SuccessResponse{})
		return
	}

	posterModel := model.Poster{}
	poster := posterModel.GetOneById(c, id)

	c.JSON(http.StatusOK, lib.SuccessResponse{
		Info: poster,
	})

}
