package cmd

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"golang.org/x/crypto/acme/autocert"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Services []Service `yaml:"services"`
}

type Service struct {
	Host    string `yaml:"host"`
	Backend string `yaml:"backend"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func NewReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}

func startProxyServer() {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	hosts := make([]string, len(config.Services))
	for i, service := range config.Services {
		hosts[i] = service.Host
	}

	m := &autocert.Manager{
		Cache:      autocert.DirCache("certs"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(hosts...),
	}

	// HTTP server to redirect HTTP to HTTPS
	httpServer := &http.Server{
		Addr: ":80",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		}),
	}

	server := &http.Server{
		Addr:      ":443",
		TLSConfig: m.TLSConfig(),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host := r.Host
			for _, service := range config.Services {
				if service.Host == host {
					proxy := NewReverseProxy(service.Backend)
					proxy.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Service not found", http.StatusNotFound)
		}),
	}

	// Start the HTTP server in a goroutine
	go func() {
		log.Println("Starting HTTP server on :80")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	log.Println("Starting proxy server on :443")
	log.Fatal(server.ListenAndServeTLS("", ""))
}
