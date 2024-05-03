package routes

import (
	"test-k-link-indonesia/handlers"
	"test-k-link-indonesia/middleware"
	"test-k-link-indonesia/packages/connection"
	"test-k-link-indonesia/repositories"

	"github.com/gin-gonic/gin"
)

func memberRoute(r *gin.RouterGroup) {
	repository := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerMember(repository)

	r.POST("/member", middleware.Auth(), handler.CreateMember)
	r.GET("/members", middleware.Auth(), handler.ShowAll)
	r.GET("/member/:id", middleware.Auth(), handler.ShowMemberByAdmin)
	r.GET("/member", middleware.Auth(), handler.ShowDataMember)

}
