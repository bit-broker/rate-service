/*
 * File Name : routers.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package routes

import (
	"net/http"

	"github.com/bit-broker/rate-service/internal/controllers"
	"github.com/gorilla/mux"
)

// CheckAPI : Returns API up
func CheckAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Rate Service API v1 Up & Running"))
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

	return router
}
