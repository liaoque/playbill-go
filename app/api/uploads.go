package api

import (
	"github.com/gin-gonic/gin"
	"image"
	"mime"
	"net/http"
	l "palybill/app/lib"
	"path"
	"strings"
)

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}

	open, err := file.Open()
	config, _, err := image.DecodeConfig(open)
	if err != nil {
		return
	}

	c.SaveUploadedFile(file, file.Filename)

	extension := strings.ToLower(path.Ext(file.Filename))
	c.JSON(http.StatusOK, l.SuccessResponse{
		Code:    0,
		Message: "",
		Info: gin.H{
			"name":      file.Filename,
			"extension": extension,
			"mime":      mime.TypeByExtension(extension),
			"size":      file.Size,
			"dimensions": map[string]int{
				"width":  config.Width,
				"height": config.Height,
			},
			"url": "1",
		},
	})

}
