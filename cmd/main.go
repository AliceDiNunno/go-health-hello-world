package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type HealthData struct {
	Running  bool
	Date     time.Time
	Hostname string
}

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.JSON(200, HealthData{
			Running:  true,
			Date:     time.Now(),
			Hostname: hostname,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
