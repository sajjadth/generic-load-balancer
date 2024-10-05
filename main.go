package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type ServerPool struct {
	servers []*url.URL
	current uint64
}

func (p *ServerPool) getNextServer() *url.URL {
	// Round-robin: get the next server
	index := atomic.AddUint64(&p.current, 1)
	return p.servers[int(index)%len(p.servers)]
}

func (p *ServerPool) loadBalancer(w http.ResponseWriter, r *http.Request) {
	target := p.getNextServer()
	proxy := httputil.NewSingleHostReverseProxy(target)

	// Rewrite the request's URL to match the target
	r.URL.Scheme = target.Scheme
	r.URL.Host = target.Host

	// Log which server is handling the request
	log.Printf("Forwarding request to: %s", target.String())

	// Forward the request
	proxy.ServeHTTP(w, r)
}

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

	var servers []*url.URL
	for _, instance := range instances {
		serverURL, err := url.Parse(instance)
		if err != nil {
			logger.Fatal("Invalid proxy instance URL", zap.String("url", instance), zap.Error(err))
		}
		servers = append(servers, serverURL)
	}

	// Create a server pool
	pool := &ServerPool{
		servers: servers,
	}

	// Start the load balancer server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default to port 8080 if not specified
	}
	// Start the load balancer
	http.HandleFunc("/", pool.loadBalancer)
	logger.Info("Load balancer server is running", zap.String("port", port))
	logger.Fatal("Server failed", zap.Error(http.ListenAndServe(":"+port, nil)))
}
