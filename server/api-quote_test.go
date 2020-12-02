package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouteGetQuote(t *testing.T) {
	// Create a test context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Test the method
	srv.RouteGetQuote(c)

	// Get the result
	assert.Equal(t, 200, w.Code, "status code is not 200")
	read, err := ioutil.ReadAll(w.Body)
	assert.NoError(t, err, "error reading response body")
	assert.NotEmpty(t, read, "response is empty")
	res := &getQuoteResponse{}
	err = json.Unmarshal(read, res)
	assert.NoError(t, err, "error parsing JSON response")
	assert.NotEmpty(t, res.Quote, "quote in the response is empty")
}
