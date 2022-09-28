package routing

import (
	"github.com/Yideg/admybrand_challenge/internal/adapter/http/rest/server"
	"github.com/gin-gonic/gin"
)

// AuthRoutes registers User Login route
func AuthRoutes(grp *gin.RouterGroup, authHandler server.AuthHandler) {
	grp.POST("/users/auth/login", authHandler.UserLogin)
}
