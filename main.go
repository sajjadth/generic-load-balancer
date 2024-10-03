package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var (
	proxyInstances  []*url.URL
	currentInstance uint64
)

func main() {
	// Initialize Zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// Load .env file if not running in Railway environment
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(filepath.Join("./", ".env")); err != nil {
			logger.Fatal("Error loading .env file", zap.Error(err))
			os.Exit(1)
		}
	}

	// Get the list of proxy instances from the environment variable
	proxyList := os.Getenv("PROXY_INSTANCES")
	if proxyList == "" {
		logger.Fatal("PROXY_INSTANCES is not set in the environment")
	}

	// Parse the proxy instances into a slice of *url.URL
	instances := strings.Split(proxyList, ",")
	for _, instance := range instances {
		parsedURL, err := url.Parse(instance)
		if err != nil {
			logger.Fatal("Invalid proxy instance URL", zap.String("url", instance), zap.Error(err))
		}
		proxyInstances = append(proxyInstances, parsedURL)
	}

	// Function to get the next proxy instance (round-robin)
	getNextProxyInstance := func() *url.URL {
		index := atomic.AddUint64(&currentInstance, 1)
		return proxyInstances[index%uint64(len(proxyInstances))]
	}

	// Define a custom transport for optimization
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// Create a custom client with the transport and set timeout
	client := &http.Client{
		Transport: transport,
	}

	// Proxy handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		targetURL := getNextProxyInstance()
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.Transport = client.Transport

		logger.Info("Proxying request: "+r.URL.String()+" + "+targetURL.String(),
			zap.String("method", r.Method),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("forwarded_to", targetURL.String()),
		)

		proxy.ServeHTTP(w, r)
	})

	// Start the load balancer server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	logger.Info("Load balancer server is running", zap.String("port", port))
	logger.Fatal("Server failed", zap.Error(http.ListenAndServe(":"+port, nil)))
}
