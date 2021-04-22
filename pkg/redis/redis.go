/*
 * File Name : mongo.go
 * Creation Date : 21-12-2020
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package redis

import (
	"github.com/bit-broker/rate-service/internal/helper"
	"strconv"

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
