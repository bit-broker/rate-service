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
 * File Name : routers.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package routes

import (
	"net/http"

	"github.com/bit-broker/rate-service/internal/controllers"
	"github.com/bit-broker/rate-service/internal/helper"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics for http duration
var (
	httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myapp_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
)

// CheckAPI : Returns API up
func CheckAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Rate Service API v1 Up & Running"))
}

// Prometheus Middleware
func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}

// InitializeRouter : Initializes the router
func InitializeRouter() *mux.Router {
	// Init Router
	router := mux.NewRouter()

	// API
	router.HandleFunc("/api/v1", CheckAPI).Methods("GET")

	// Rate Service
	router.Handle("/api/v1/{uid}/config", http.HandlerFunc(controllers.GetConfig)).Methods("GET")
	router.Handle("/api/v1/{uid}/config", http.HandlerFunc(controllers.CreateOrUpdateConfig)).Methods("PUT")
	router.Handle("/api/v1/{uid}/config", http.HandlerFunc(controllers.DeleteConfig)).Methods("DELETE")

	// Metrics
	if helper.GetConfiguration().MetricsEnabled == "true" {
		router.Use(prometheusMiddleware)
		router.Path("/metrics").Handler(promhttp.Handler())
	}

	return router
}
