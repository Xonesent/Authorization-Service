package handler

import (
	handler_api "project/pkg/handler/api"
	"project/pkg/service"

	"github.com/gin-gonic/gin"
)

type Authorization interface{
	Sign_Up(c *gin.Context)
	Sign_In(c *gin.Context)
	User_Identity(c *gin.Context)
}

type Links interface{
	Send_Url(c *gin.Context)
	Get_Url(c *gin.Context)
}

type Handler struct {
	Authorization
	Links
}

func New_Handler(services *service.Service) *Handler {
	return &Handler{
		Authorization: handler_api.New_Auth_handler(services.Authorization),
		Links: handler_api.New_Links_handler(services.Links),
	}
}

func (h *Handler) Init_Routes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.Sign_Up)
		auth.POST("/sign-in", h.Sign_In)
	}

	api := router.Group("/api", h.User_Identity)
	{
		api.POST("send_url", h.Send_Url)
		api.GET("get_url/:short_url", h.Get_Url)
	}

	return router
}
