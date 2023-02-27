package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-devops/golang-oauth2-example/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Info struct {
	UpTime time.Duration `json:"up_time"`
	Status string        `json:"status" `
}

var (
	startTime time.Time
)

func Init() {
	startTime = time.Now()
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func getInfos(c *gin.Context) {
	state := Info{
		UpTime: uptime(),
		Status: "ok",
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, state)
}

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	priv := r.Group("/priv")           // private endpoints
	priv.Use(auth.JwtAuthMiddleware()) // private endpoints protection

	priv.GET("/info", getInfos)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("service listen on port %s\n", "9100")
	log.Fatal(srv.ListenAndServe())
}
