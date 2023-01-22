package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project_dwi/helper"
)

type helloHandler struct{}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h *helloHandler) SayHello(c *gin.Context) {
	response := helper.APIResponse("Hi, This is Ethereal Beauty Backend App", http.StatusOK, "Success", nil)
	c.JSON(http.StatusOK, response)
}
