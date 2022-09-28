package routing

import (
	"github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server"
	"github.com/gin-gonic/gin"
)

// UserRoutes registers user list routes
func UserRoutes(grp *gin.RouterGroup, userHandler server.UserHandler) {
	grp.GET("/users", userHandler.UsersHandler)
	grp.GET("/users/:id", userHandler.GetUserByIDHandler)
	grp.POST("/users", userHandler.StoreUserHandler)
	grp.PUT("/users/:id", userHandler.UpdateUserHandler)
	grp.PATCH("/users/assign-role", userHandler.AssignRoleToUser)
	grp.DELETE("/users/:id", userHandler.DeleteUserHandler)

}
