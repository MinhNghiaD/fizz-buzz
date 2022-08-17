package main

import (
	"fmt"

	"github.com/MinhNghiaD/fizz-buzz/pkg/handler"
	"github.com/gin-gonic/gin"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port = kingpin.Flag("port", "server port").Short('p').Default("8080").Int()
)

func main() {
	kingpin.Parse()
	router := gin.Default()
	fizzbuzzHandler := handler.NewFizzbuzzHandler()

	router.POST("/fizzbuzz", fizzbuzzHandler.HandleFizzbuzz)
	router.GET("/most-frequent", fizzbuzzHandler.HandleMostFrequent)

	err := router.Run(fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}
}
