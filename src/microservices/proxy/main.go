package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

func main() {
	port := getEnv("PORT", "8000")
	monolithURL := getEnv("MONOLITH_URL", "http://localhost:8080")
	moviesServiceURL := getEnv("MOVIES_SERVICE_URL", "http://localhost:8081")

	gradualMigration := getEnvBool("GRADUAL_MIGRATION", false)
	moviesMigrationPercent := getEnvInt("MOVIES_MIGRATION_PERCENT", 0)

	monolithProxy := createReverseProxy(monolithURL)
	moviesProxy := createReverseProxy(moviesServiceURL)

	http.HandleFunc("/api/movies", func(w http.ResponseWriter, r *http.Request) {
		if shouldRouteToMoviesService(gradualMigration, moviesMigrationPercent) {
			log.Printf(
				"Routing request to movies-service, migration percent: %d%%",
				moviesMigrationPercent,
			)

			moviesProxy.ServeHTTP(w, r)
			return
		}

		log.Printf(
			"Routing request to monolith, migration percent: %d%%",
			moviesMigrationPercent,
		)

		monolithProxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Routing request to monolith")
		monolithProxy.ServeHTTP(w, r)
	})

	log.Printf(
		"Starting proxy service on port %s, gradual migration: %t, movies migration percent: %d%%",
		port,
		gradualMigration,
		moviesMigrationPercent,
	)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func shouldRouteToMoviesService(
	gradualMigration bool,
	migrationPercent int,
) bool {
	if !gradualMigration {
		return false
	}

	randomValue := rand.Intn(100)

	return randomValue < migrationPercent
}

func createReverseProxy(target string) *httputil.ReverseProxy {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Invalid target URL %s: %v", target, err)
	}

	return httputil.NewSingleHostReverseProxy(targetURL)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsedValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatalf("Invalid boolean value for %s: %s", key, value)
	}

	return parsedValue
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Invalid integer value for %s: %s", key, value)
	}

	if parsedValue < 0 || parsedValue > 100 {
		log.Fatalf("%s must be between 0 and 100", key)
	}

	return parsedValue
}