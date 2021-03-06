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
 * File Name : rate.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bit-broker/rate-service/internal/helper"
	"github.com/bit-broker/rate-service/internal/models"
	"github.com/bit-broker/rate-service/internal/services"
	"github.com/bit-broker/rate-service/pkg/log"

	ratelimit "github.com/datawire/ambassador/pkg/api/envoy/service/ratelimit/v2"

	"github.com/gorilla/mux"
)

// ------------------------ HTTP REST -------------------- //

// GetConfig : CRUD
func GetConfig(w http.ResponseWriter, r *http.Request) {
	log.Info("Returning config")

	// Get params
	var params = mux.Vars(r)
	uid := params["uid"]

	// Get config
	config, err := services.GetConfig(uid)

	if err != nil {
		helper.GetNotFoundError(w)
		return
	}

	// Remove log
	config.Log = nil

	// Set header.
	w.Header().Set("Content-Type", "application/json")

	// Response
	_ = json.NewEncoder(w).Encode(config)
}

// CreateOrUpdateConfig : CRUD
func CreateOrUpdateConfig(w http.ResponseWriter, r *http.Request) {
	log.Info("Creating or Updating config")

	// Get params
	var params = mux.Vars(r)
	uid := params["uid"]

	// Decode body
	var config models.Config
	_ = json.NewDecoder(r.Body).Decode(&config)

	// Create config
	err := services.CreateOrUpdateConfig(uid, config)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Set header.
	w.Header().Set("Content-Type", "application/json")

	// Response
	_ = json.NewEncoder(w).Encode(config)
}

// DeleteConfig : CRUD
func DeleteConfig(w http.ResponseWriter, r *http.Request) {
	log.Info("Deleting config")

	// Get params
	var params = mux.Vars(r)
	uid := params["uid"]

	// Delete config
	err := services.DeleteConfig(uid)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Set header.
	w.Header().Set("Content-Type", "application/json")

	// Response
	_ = json.NewEncoder(w).Encode("OK")
}

// ------------------------ HTTP REST -------------------- //

// ------------------------ GRPC ------------------------- //

// RatelimitService : gRPC Rate Limit Interface
type RatelimitService struct {
}

// ShouldRateLimit : gRPC Rate Limit Interface
func (r RatelimitService) ShouldRateLimit(ctx context.Context, request *ratelimit.RateLimitRequest) (*ratelimit.RateLimitResponse, error) {
	log.Info("Received request", request)

	// Get uid
	var uid string
	for _, descriptor := range request.Descriptors {
		for _, entry := range descriptor.Entries {
			if entry.Key == "uid" {
				uid = entry.Value
			}
		}
	}

	// If uid not present
	if len(uid) <= 0 {
		log.Debug("Over Limit")
		return &ratelimit.RateLimitResponse{
			OverallCode: ratelimit.RateLimitResponse_OVER_LIMIT,
		}, nil
	}

	// Check config
	ok, _ := services.Check(uid)

	if !ok {
		log.Debug("Over Limit")
		return &ratelimit.RateLimitResponse{
			OverallCode: ratelimit.RateLimitResponse_OVER_LIMIT,
		}, nil
	}

	log.Debug("OK")
	return &ratelimit.RateLimitResponse{
		OverallCode: ratelimit.RateLimitResponse_OK,
	}, nil
}

// ------------------------ GRPC ------------------------- //
