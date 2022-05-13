// Copyright 2021 Cisco and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

/*
 * File Name : configuration.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
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
