/*
 * File Name : main.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
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
	ratelimit_v1 "github.com/datawire/ambassador/pkg/api/pb/lyft/ratelimit"

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
	go startGRPCServer(config.ServerGRPCHost)

	// Start HTTP Server
	handler := c.Handler(router)
	log.Info("Starting HTTP Server")
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
	ratelimit_v1.RegisterRateLimitServiceServer(gRPCServer, controllers.RatelimitService{})
	ratelimit_v2.RegisterRateLimitServiceServer(gRPCServer, controllers.RatelimitService{})
	log.Info("Starting gRPC Server")
	if err := gRPCServer.Serve(listner); err != nil {
		log.Fatal("gRPC failed", err)
	}
}
