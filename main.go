package main

import (
	"fmt"
	"harmony-market/graphql"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Print("[server] start graphql server...")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://www.harmony-market.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	gqlHandler, err := graphql.CreateGraphQLHandler()
	if err != nil {
		log.Fatal(err)
	}

	r.Any("/graphql", func(c *gin.Context) {
		gqlHandler.ServeHTTP(c.Writer, c.Request)
	})

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
