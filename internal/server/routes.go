package server

import (
	"go-prototype-new/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(healtcheckController controller.HomeController) http.Handler {
	r := gin.Default()

	r.GET("/", healtcheckController.GetHome)

	r.GET("/health", healtcheckController.GetHealthCheck)

	return r
}
