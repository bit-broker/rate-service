/*
 * File Name : configuration.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package helper

import (
	"os"

	"github.com/joho/godotenv"
)

// Configuration model.
type Configuration struct {
	ServerHTTPHost             string
	ServerGRPCHost             string
	GoEnv                      string
	UIUrl                      string
	LogLevel                   string
	RedisAddr                  string
	RedisPassword              string
	RedisDB                    string
	PolicyServiceEndpoint      string
	PolicyServiceAuthorization string
	PolicyServiceTimeout       string
	MetricsEnabled             string
}

// Env : Type of env
type Env string

// Env types
const (
	TestEnv Env = "../../../.env.test"
	AllEnv  Env = ".env"
)

// LoadEnv : Load env
func LoadEnv(env Env) {
	_ = godotenv.Load(string(env))
}

// GetConfiguration : Populate from env
func GetConfiguration() Configuration {
	configuration := Configuration{
		ServerHTTPHost:             os.Getenv("SERVER_HTTP_HOST"),
		ServerGRPCHost:             os.Getenv("SERVER_GRPC_HOST"),
		GoEnv:                      os.Getenv("GO_ENV"),
		UIUrl:                      os.Getenv("UI_URL"),
		LogLevel:                   os.Getenv("LOG_LEVEL"),
		RedisAddr:                  os.Getenv("REDIS_ADDR"),
		RedisPassword:              os.Getenv("REDIS_PASSWORD"),
		RedisDB:                    os.Getenv("REDIS_DB"),
		PolicyServiceEndpoint:      os.Getenv("POLICY_SERVICE_ENDPOINT"),
		PolicyServiceAuthorization: os.Getenv("POLICY_SERVICE_AUTHORIZATION"),
		PolicyServiceTimeout:       os.Getenv("POLICY_SERVICE_TIMEOUT"),
		MetricsEnabled:             os.Getenv("METRICS_ENABLED"),
	}

	return configuration
}
