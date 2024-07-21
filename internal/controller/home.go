package controller

import (
	"go-prototype-new/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	HomeController interface {
		GetHome(c *gin.Context)
		GetHealthCheck(c *gin.Context)
	}

	homeController struct {
		Repository repository.HomeRepository
	}
)

func NewHomeController(homeRepository repository.HomeRepository) HomeController {
	return &homeController{
		Repository: homeRepository,
	}
}

func (controller *homeController) GetHome(c *gin.Context) {
	resp := make(map[string]interface{})

	//call model
	home := controller.Repository.GetHome(c.Request.Context())

	resp["data"] = home

	c.JSON(http.StatusOK, resp)
}

func (controller *homeController) GetHealthCheck(c *gin.Context) {
	resp := make(map[string]string)

	resp["message"] = "Health Check is Okay.."

	c.JSON(http.StatusOK, resp)
}
