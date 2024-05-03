package routes

import (
	"test-k-link-indonesia/handlers"
	"test-k-link-indonesia/packages/connection"
	"test-k-link-indonesia/repositories"

	"github.com/gin-gonic/gin"
)

func userRoute(r *gin.RouterGroup) {
	repository := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerUser(repository)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

}
