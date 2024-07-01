package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"go-prototype-new/internal/app/database"
	"go-prototype-new/internal/controller"
	"go-prototype-new/internal/model"
)

func NewServer() *http.Server {
	db := database.New()

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	homeModel := model.NewHomeModel(db)
	homeController := controller.NewHomeController(homeModel)

	router := RegisterRoutes(homeController)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
