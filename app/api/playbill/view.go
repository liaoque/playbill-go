package playbill

import (
	"github.com/gin-gonic/gin"
	"net/http"
	l "palybill/app/lib"
	"palybill/app/model"
)

func viewHandler(c *gin.Context) {
	id := c.Param("id")
	//base64 := c.Param("base64")
	if len(id) > 0 {
		//$params = $this->getRequest()->getParams();
		//if (empty($params)) {
		//throw new Yaf_Exception('参数不能为空', \AppResponse\AppResponsePlayBill::PARAMS_EMPTY);
		//}
		//$params = json_decode(json_encode($params));
		//$out = \PlayBill\Factory::load($params);
	} else {
		posterModel := model.Poster{}
		poster := posterModel.GetOneById(c, id)
		if len(poster.Id) == 0 {
			c.JSON(http.StatusBadRequest, l.ErrorResponse{
				Code:    101,
				Message: "id not found",
			})
			return
		}
		print(poster)
		//$params = $this->getRequest()->getParams();
		//$poster = new PosterModel();
		//$data = $poster->getRowById($id);
		//$out = \PlayBill\Factory::load($data, $params);
		//$params = $data;
	}

}
