package handler_api

import (
	"net/http"
	"project/pkg/service"
	"project/server"

	"github.com/gin-gonic/gin"
)

type Links_handler struct {
	links service.Links
}

func New_Links_handler(links service.Links) *Links_handler {
	return &Links_handler{links: links}
}

func (h *Links_handler) Send_Url(c *gin.Context) {
	var input server.Link

	if err := c.BindJSON(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "invalid input body error")
		return
	}

	if err := server.Validate_Base_URL(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "validation error")
		return
	}

	short_url, err := h.links.Create_Short_URL(&input)
	if err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "impossible to create short url")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Short_URL": short_url,
	})
}

func (h *Links_handler) Get_Url(c *gin.Context) {
	var input server.Link
	input.Short_URL = c.Param("short_url")

	if err := server.Validate_Short_URL(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "validation error")
		return
	}

	base_url, err := h.links.Get_Base_URL(&input)
	if err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "impossible to get base url")
		return
	}

	if base_url == "" {
		server.New_Error_Response(c, http.StatusBadRequest, "empty created short url")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Base_URL": base_url,
	})
}
