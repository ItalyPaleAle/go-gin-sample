package server

import (
	"flag"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var srv *Server

func TestMain(m *testing.M) {
	// Parse flags
	flag.Parse()

	// Init the server
	srv = &Server{}
	err := srv.Init()
	if err != nil {
		panic(err)
	}

	// Set Gin to testing mode
	gin.SetMode(gin.TestMode)

	// Run tests
	os.Exit(m.Run())
}
