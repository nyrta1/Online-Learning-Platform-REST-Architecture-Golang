package routers

import (
	"github.com/gin-gonic/gin"
	"online-learning-platform/internal/rest/handlers"
)

type Routers struct {
	authHandlers handlers.AuthHandlers
}

func NewRouters(authHandlers handlers.AuthHandlers) *Routers {
	return &Routers{
		authHandlers: authHandlers,
	}
}

func (r *Routers) SetupRoutes(app *gin.Engine) {
	authRouter := app.Group("/auth")
	{
		authRouter.POST("/register", r.authHandlers.Register)
		//authRouter.POST("/login", r.authHandlers.Login)
		//authRouter.POST("/logout", middleware.RequireAuthMiddleware, r.authHandlers.Logout)
		//authRouter.DELETE("/deleteAccount", middleware.RequireAuthMiddleware, r.authHandlers.DeleteAccount)
	}
}
