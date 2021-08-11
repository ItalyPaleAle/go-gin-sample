package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
)

// Server is the server based on Gin
type Server struct {
	router     *gin.Engine
	httpClient *http.Client
}

// Init the Server object and create a Gin server
func (s *Server) Init() error {
	// Set Gin to Release mode
	gin.SetMode(gin.ReleaseMode)

	// Init a HTTP client
	s.httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create the Gin router and add the Recovery middleware to recover from panics
	s.router = gin.New()
	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())

	// Enable CORS from any origin
	s.router.Use(cors.Default())

	// Add REST routes
	s.router.GET("/api/quote2", s.RouteGetQuote)

	// Serve the frontend application
	err := s.serveFrontend()
	if err != nil {
		return err
	}

	return nil
}

// Start the web server
// Note this function is blocking, and will return only when the servers are shut down (via context cancelation or via SIGINT/SIGTERM signals)
func (s *Server) Start(ctx context.Context) {
	// Get address and ports from env vars or fallback to defaults
	bindAddr := os.Getenv("BIND")
	if bindAddr == "" {
		bindAddr = "127.0.0.1"
	}
	httpPort, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if httpPort == 0 {
		httpPort = 8080
	}
	httpsPort, _ := strconv.Atoi(os.Getenv("HTTPS_PORT"))
	if httpsPort == 0 {
		httpsPort = 8443
	}

	// TLS certificate and key
	var tlsCert, tlsKey string
	if os.Getenv("NO_TLS") != "1" {
		tlsCert = os.Getenv("TLS_CERT")
		if tlsCert == "" {
			tlsCert = "certs/cert.pem"
		}
		tlsKey = os.Getenv("TLS_KEY")
		if tlsKey == "" {
			tlsKey = "certs/key.pem"
		}
	}

	// Launch the server (this is a blocking call)
	s.launchServer(ctx, bindAddr, httpPort, httpsPort, tlsCert, tlsKey)
}

// Start the server
func (s *Server) launchServer(ctx context.Context, bindAddr string, httpPort, httpsPort int, tlsCert, tlsKey string) {
	// If we don't have a TLS certificate, don't enable TLS
	enableTLS := (tlsCert != "" && tlsKey != "")

	// HTTP server (no TLS)
	httpSrv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", bindAddr, httpPort),
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// HTTPS server (with TLS)
	httpsSrv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", bindAddr, httpsPort),
		Handler:        s.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Start the HTTP server in a background goroutine
	go func() {
		fmt.Printf("HTTP server listening on http://%s:%d\n", bindAddr, httpPort)
		// Next call blocks until the server is shut down
		err := httpSrv.ListenAndServe()
		if err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Start the HTTPS server in a background goroutine
	if enableTLS {
		go func() {
			fmt.Printf("HTTPS server listening on https://%s:%d\n", bindAddr, httpsPort)
			err := httpsSrv.ListenAndServeTLS(tlsCert, tlsKey)
			if err != http.ErrServerClosed {
				panic(err)
			}
		}()
	}

	// Listen to SIGINT and SIGTERM signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	// Block until we either get a termination signal, or until the context is canceled
	select {
	case <-ctx.Done():
	case <-ch:
	}

	// We received an interrupt signal, shut down both servers
	var errHttp, errHttps error
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	errHttp = httpSrv.Shutdown(shutdownCtx)
	if enableTLS {
		errHttps = httpsSrv.Shutdown(shutdownCtx)
	}
	shutdownCancel()
	// Log the errors (could be context canceled)
	if errHttp != nil {
		log.Println("HTTP server shutdown error:", errHttp)
	}
	if errHttps != nil {
		log.Println("HTTPS server shutdown error:", errHttps)
	}
}

func (s *Server) serveFrontend() error {
	// Serve the (compiled) frontend app when requesting a path that doesn't exist
	box, err := pkger.Open("/frontend/public")
	if err != nil {
		return err
	}
	s.router.NoRoute(gin.WrapH(http.FileServer(box)))

	return nil
}
