package handler

import (
	"fmt"
	"net/http"

	"github.com/MinhNghiaD/fizz-buzz/pkg/fizzbuzz"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// fizzbuzzRequest encapsulates a fizzbuzz json request body
type fizzbuzzRequest struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// FizzbuzzHandler handles fizzbuzz requests
type FizzbuzzHandler struct {
	requestCounter fizzbuzz.RequestCounter
}

// NewFizzbuzzHandler initiates new instance of FizzbuzzHandler
func NewFizzbuzzHandler() *FizzbuzzHandler {
	return &FizzbuzzHandler{
		requestCounter: fizzbuzz.NewRequestCounter(),
	}
}

// HandleFizzbuzz handles fizzbuzz request. A list of string correspond to the request is returned. If the arguments provided by the request body is bad, an error will be returned instead
func (handler *FizzbuzzHandler) HandleFizzbuzz(c *gin.Context) {
	var request fizzbuzzRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	if err := verifyRequest(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handler.requestCounter.NewRequest(request)
	respond := fizzbuzz.FizzBuzz(request.Int1, request.Int2, request.Limit, request.Str1, request.Str2)
	log.Debugf("Respond %v for request %v", respond, request)

	c.JSON(http.StatusOK, respond)
}

// HandleMostFrequent handles most-frequent request by returning the most frequent requests handling by the server as well as the number of these requests
func (handler *FizzbuzzHandler) HandleMostFrequent(c *gin.Context) {
	requests, frequency := handler.requestCounter.MostFrequentRequest()
	if frequency < 0 {
		c.IndentedJSON(http.StatusOK, gin.H{})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"requests": requests, "count": frequency})
}

// verifyRequest verifies the arguments of a fizzbuzz request
func verifyRequest(request fizzbuzzRequest) error {
	if request.Int1 <= 0 || request.Int2 <= 0 || request.Limit <= 0 || request.Str1 == "" || request.Str2 == "" {
		return fmt.Errorf("missing arguments")
	}

	return nil
}
