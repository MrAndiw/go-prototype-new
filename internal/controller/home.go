package controller

import (
	"go-prototype-new/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	HomeController interface {
		GetHome(c *gin.Context)
		GetHealthCheck(c *gin.Context)
	}

	homeController struct {
		model model.HomeModel
	}
)

func NewHomeController(model model.HomeModel) HomeController {
	return &homeController{
		model: model,
	}
}

func (controller *homeController) GetHome(c *gin.Context) {
	resp := make(map[string]string)

	//call model
	home := controller.model.GetHome("Hello World")

	resp["message"] = home

	c.JSON(http.StatusOK, resp)
}

func (controller *homeController) GetHealthCheck(c *gin.Context) {
	resp := make(map[string]string)

	resp["message"] = "Health Check is Okay.."

	c.JSON(http.StatusOK, resp)
}
