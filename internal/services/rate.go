/*
 * File Name : rate.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/bit-broker/rate-service/internal/helper"
	"github.com/bit-broker/rate-service/internal/models"

	"github.com/bit-broker/rate-service/pkg/redis"
)

// ------------------------ GLOBAL -------------------- //

var redisContext = context.Background()

const dayLayout = "%d-%02d-%02d"
const monthLayout = "%d-%02d"
const defaultTimeout = 5

// ------------------------ GLOBAL -------------------- //

// GetConfig : CRUD
func GetConfig(uid string) (models.Config, error) {
	// Get config
	var config models.Config
	raw, err := redis.Client().Get(redisContext, uid).Result()
	json.Unmarshal([]byte(raw), &config)

	return config, err
}

// CreateOrUpdateConfig : CRUD
func CreateOrUpdateConfig(uid string, config models.Config) error {
	// Set config
	raw, _ := json.Marshal(config)
	_, err := redis.Client().Set(redisContext, uid, raw, 0).Result()

	return err
}

// DeleteConfig : CRUD
func DeleteConfig(uid string) error {
	// Delete config
	_, err := redis.Client().Del(redisContext, uid).Result()

	return err
}

// FetchConfig : If config cannot be found locally, fallback
// to the policy service hook
func FetchConfig(uid string) (models.Config, error) {
	// Get config
	var config models.Config

	// Check endpoint definition
	if len(helper.GetConfiguration().PolicyServiceEndpoint) <= 0 {
		return config, errors.New("Policy Service Endpoint not defined")
	}

	// Get request timeout
	timeout := defaultTimeout
	if len(helper.GetConfiguration().PolicyServiceTimeout) > 0 {
		timeout, _ = strconv.Atoi(helper.GetConfiguration().PolicyServiceTimeout)
	}

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// Add endpoint URL
	req, _ := http.NewRequestWithContext(ctx, "GET", helper.GetConfiguration().PolicyServiceEndpoint, nil)

	// Add authorization header if specified
	if authorization := helper.GetConfiguration().PolicyServiceAuthorization; len(authorization) > 0 {
		req.Header.Add("Authorization", authorization)
	}

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return config, err
	}

	// Read and parse data
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal([]byte(body), &config)
	_ = resp.Body.Close()

	return config, err
}

// Check : Check if current request is within the config
func Check(uid string) (bool, error) {
	config, err := GetConfig(uid)

	// Check err
	if err != nil {
		// Try to fallback on the policy service endpoint
		config, err = FetchConfig(uid)

		if err != nil {
			return false, err
		}

		// Cache config
		go CreateOrUpdateConfig(uid, config)
	}

	// Check if enabled
	if !config.Enabled {
		return false, nil
	}

	// Create if needed
	if len(config.Log) <= 0 {
		config.Log = make(map[string]int)
	}

	// Current time
	currentTime := time.Now()

	// Save config
	defer CreateOrUpdateConfig(uid, config)

	// Rate / s
	// Get current time
	withinRate := true
	now := strconv.FormatInt(currentTime.Unix(), 10)
	if val, ok := config.Log[now]; ok {
		// Check rate
		if val >= config.Rate {
			withinRate = false
		} else {
			config.Log[now] = val + 1
		}
	} else {
		config.Log[now] = 1
	}

	if !withinRate {
		return false, nil
	}

	// Quota / Interval
	// Get interval type
	withinQuota := true
	var interval string
	switch config.Quota.Interval {
	case models.DayType:
		interval = fmt.Sprintf(dayLayout,
			currentTime.Year(), currentTime.Month(), currentTime.Day())
	default: // Default is month
		interval = fmt.Sprintf(monthLayout,
			currentTime.Year(), currentTime.Month())
	}

	if val, ok := config.Log[interval]; ok {
		// Check quota
		if val >= config.Quota.Number {
			withinQuota = false
		} else {
			config.Log[interval] = val + 1
		}
	} else {
		config.Log[interval] = 1
	}

	return withinRate && withinQuota, nil
}
