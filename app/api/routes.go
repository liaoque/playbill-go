package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routesHandler(c *gin.Context) {
	//	return AppResponse\AppResponse::success([
	//		'token' => 'eyJhbGciOiJIUzUxMiJ9.test'
	//]);
	menu := `{"path":"/permission","meta":{"title":"menus.permission","icon":"lollipop","rank":7},"children":[{"path":"/permission/page/index","name":"PermissionPage","meta":{"title":"menus.permissionPage"}},{"path":"/permission/button/index","name":"PermissionButton","meta":{"title":"menus.permissionButton","authority":[]}}]}`

	c.String(http.StatusOK, menu)
	//c.String()
}
