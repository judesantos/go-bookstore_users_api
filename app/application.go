package app

// application.go

import (
	"github.com/gin-gonic/gin"
	"github.com/judesantos/go-bookstore_utils/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("Starting application at port 8080...")
	router.Run(":8080")
}
