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
 * File Name : main.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/rs/cors"

	"github.com/bit-broker/rate-service/internal/controllers"
	"github.com/bit-broker/rate-service/internal/helper"
	"github.com/bit-broker/rate-service/internal/routes"
	"github.com/bit-broker/rate-service/pkg/log"

	ratelimit_v2 "github.com/datawire/ambassador/pkg/api/envoy/service/ratelimit/v2"

	"google.golang.org/grpc"
)

func main() {
	// Load env
	helper.LoadEnv(helper.AllEnv)

	router := routes.InitializeRouter()
	config := helper.GetConfiguration()
	log.Info("Starting in env ", config.GoEnv)

	// Configure log level
	log.SetLogLevel(config.LogLevel)

	// Setup cors for dev
	options := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "GET", "DELETE", "POST", "PATCH"},
		AllowedHeaders:   []string{"X-Requested-With", "content-type", "Origin", "Accept", "Authorization"},
		AllowCredentials: true,

		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	}

	// Check current env
	if config.GoEnv != "development" {
		options.AllowedOrigins = strings.Split(config.UIUrl, ",")
		options.Debug = false
	}
	c := cors.New(options)

	// Start GRPC Server
	log.Info("Starting GRPC Server with ", config.ServerGRPCHost)
	go startGRPCServer(config.ServerGRPCHost)

	// Start HTTP Server
	handler := c.Handler(router)
	log.Info("Starting HTTP Server with ", config.ServerHTTPHost)
	log.Fatal(http.ListenAndServe(config.ServerHTTPHost, handler))
}

// gRPC Server
func startGRPCServer(host string) {
	listner, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal("gRPC failed", err)
	}
	gRPCServer := grpc.NewServer()

	// Register the service
	ratelimit_v2.RegisterRateLimitServiceServer(gRPCServer, controllers.RatelimitService{})
	log.Info("Starting gRPC Server")
	if err := gRPCServer.Serve(listner); err != nil {
		log.Fatal("gRPC failed", err)
	}
}
