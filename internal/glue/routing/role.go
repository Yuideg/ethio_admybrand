package routing

import (
	"github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server"
	"github.com/gin-gonic/gin"
)

// RoleRoutes registers Role of  list routes
func RoleRoutes(grp *gin.RouterGroup, roleHandler server.RoleHandler) {
	grp.GET("/roles", roleHandler.RolesHandler)
	grp.POST("/roles", roleHandler.StoreRoleHandler)
	grp.DELETE("/roles", roleHandler.DeleteRoleHandler)
}
