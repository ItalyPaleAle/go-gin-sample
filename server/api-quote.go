package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type taylorRestResponse struct {
	Quote string `json:"quote"`
}

type getQuoteResponse struct {
	Quote string `json:"quote"`
}

// RouteGetQuote is the handler for the GET /api/quote request
// The response contains a random Taylor Swift quote, as served by https://api.taylor.rest/
func (s *Server) RouteGetQuote(c *gin.Context) {
	// Request a new quote from the external service
	res, err := s.httpClient.Get("https://api.taylor.rest/")
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to request a quote from the external service",
		})
		return
	}
	defer res.Body.Close()

	// Read the entire response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": "Error while reading the respone from the external service",
		})
		return
	}
	if len(body) == 0 {
		c.Error(errors.New("empty response from the external service"))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": "Received an empty respone from the external service",
		})
		return
	}

	// Parse the response
	quote := &taylorRestResponse{}
	err = json.Unmarshal(body, quote)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": "Error while parsing the JSON respone from the external service",
		})
		return
	}
	if quote.Quote == "" {
		c.Error(errors.New("response from the external service does not contain a quote"))
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"error": "The respone from the external service does not contain a quote",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, &getQuoteResponse{
		Quote: quote.Quote,
	})
}
