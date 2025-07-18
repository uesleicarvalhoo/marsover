package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var config = map[string]string{
	// Application
	"SERVICE_NAME":    "foxbit-test",
	"SERVICE_VERSION": "0.0.0",
	"LOG_LEVEL":       "INFO",
	"ENVIRONMENT":     "dev",

	// Orchestrator
	"ORCHESTRATOR_CONCURRENCY": "10",

	// HTTP Server
	"HTTP_PORT":                 "8080",
	"HTTP_SERVER_READ_TIMEOUT":  "15s",
	"HTTP_SERVER_WRITE_TIMEOUT": "15s",
}

// GetString value of a given env var
func GetString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		return config[k]
	}

	return v
}

// GetInt value of a given env var
func GetInt(k string) int {
	v := GetString(k)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

// Get value of a given env var
func GetFloat64(k string) float64 {
	v := GetString(k)
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}

	return f
}

// GetDuration value of a given env var
func GetDuration(k string) time.Duration {
	d, err := time.ParseDuration(GetString(k))
	if err != nil {
		panic(err)
	}

	return d
}

// GetBool value of a given env var
func GetBool(k string) bool {
	v := GetString(k)

	return strings.ToLower(v) == "true"
}

// Set config for test purposes
func Set(k, v string) {
	config[k] = v
}
