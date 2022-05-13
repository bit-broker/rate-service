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
 * File Name : mongo.go
 * Creation Date : 21-12-2020
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package redis

import (
	"strconv"

	"github.com/bit-broker/rate-service/internal/helper"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

// ------------------------ GLOBAL -------------------- //

var redisClient *redis.Client = nil

// ------------------------ GLOBAL -------------------- //

// Client : This is a helper function to connect to Redis
func Client() *redis.Client {
	// Get real config
	config := helper.GetConfiguration()

	// Check invalid config
	if len(config.RedisAddr) <= 0 {
		return nil
	}

	// Check already initialized
	if redisClient != nil {
		return redisClient
	}

	// Check if test env
	if config.RedisAddr == "mockup" {
		redisServer := mockRedis()

		redisClient = redis.NewClient(&redis.Options{
			Addr: redisServer.Addr(),
		})

		return redisClient
	}

	// Set client options
	db, _ := strconv.Atoi(config.RedisDB)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       db,
	})

	return redisClient
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}
