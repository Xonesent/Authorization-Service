package handler_api

import (
	"net/http"
	"strings"
	"project/pkg/service"
	"project/server"

	"github.com/gin-gonic/gin"
)

type Register_Input struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Auth_handler struct {
	authorization service.Authorization
}

func New_Auth_handler(authorization service.Authorization) *Auth_handler {
	return &Auth_handler{authorization: authorization}
}

func (h *Auth_handler) Sign_Up(c *gin.Context) {
	var input Register_Input

	if err := c.BindJSON(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.authorization.Create_User(input.Login, input.Password)
	if err != nil {
		server.New_Error_Response(c, http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Auth_handler) Sign_In(c *gin.Context) {
	var input Register_Input

	if err := c.BindJSON(&input); err != nil {
		server.New_Error_Response(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.authorization.Generate_Token(input.Login, input.Password)
	if err != nil {
		server.New_Error_Response(c, http.StatusInternalServerError, "server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Auth_handler) User_Identity(c *gin.Context) {
	Header := c.GetHeader("Authorization")
	if Header == "" {
		server.New_Error_Response(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	Header_Parts := strings.Split(Header, " ")
	if len(Header_Parts) != 2 || Header_Parts[0] != "Bearer" {
		server.New_Error_Response(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(Header_Parts[1]) == 0 {
		server.New_Error_Response(c, http.StatusUnauthorized, "token is empty")
		return
	}

	User_Id, err := h.authorization.Parse_Token(Header_Parts[1])
	if err != nil {
		server.New_Error_Response(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("user_id", User_Id)
}
