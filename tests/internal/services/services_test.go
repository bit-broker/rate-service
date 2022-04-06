/*
 * File Name : services_test.go
 * Creation Date : 11-05-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package tests

import (
	"context"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/bit-broker/rate-service/internal/helper"
	"github.com/bit-broker/rate-service/internal/models"
	"github.com/bit-broker/rate-service/internal/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// ------------------------ GLOBAL -------------------- //

var uid = strconv.Itoa(rand.Intn(100))
var mockupFirstConfigRaw = `{"enabled":true,"quota":{"max_number":10,"interval_type":"month"},"rate":5}`
var mockupFirstConfig = &models.Config{Enabled: true, Quota: models.Quota{Number: 10, Interval: models.MonthType}, Rate: 5}
var mockupSecondConfig = &models.Config{Enabled: true, Quota: models.Quota{Number: 5, Interval: models.DayType}, Rate: 20}
var mockupPolicyServerAddr = ":5001"

// ------------------------ GLOBAL -------------------- //

// TestServices : Services Test cases
func TestControllers(t *testing.T) {
	// Load env
	helper.LoadEnv(helper.TestEnv)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Test Suite")
}

var _ = Describe("Services", func() {
	Context("Get", func() {
		It("shouldn't find the config when not created", func() {
			// Get config
			_, err := services.GetConfig(uid)
			Expect(err).NotTo(BeNil())
		})

		It("should find the config after it was created", func() {
			// Create config
			services.CreateOrUpdateConfig(uid, *mockupFirstConfig)

			// Get config
			config, err := services.GetConfig(uid)
			Expect(err).To(BeNil())
			Expect(config).To(Equal(*mockupFirstConfig))
		})

		It("should find it when the policy server is defined", func() {
			// Delete config
			services.DeleteConfig(uid)

			// Get config
			_, err := services.GetConfig(uid)
			Expect(err).NotTo(BeNil())

			// Start server
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				_, _ = w.Write([]byte(mockupFirstConfigRaw))
			})
			mockupPolicyServer := &http.Server{Addr: mockupPolicyServerAddr}
			go func() {
				mockupPolicyServer.ListenAndServe()
			}()

			// Try to get config
			status, err := services.Check(uid)
			Expect(err).To(BeNil())
			Expect(status).To(BeTrue())

			// Check current config
			config, err := services.GetConfig(uid)
			Expect(err).To(BeNil())
			config.Log = nil // Clear the log as check was called
			Expect(config).To(Equal(*mockupFirstConfig))

			// Close server
			mockupPolicyServer.Shutdown(context.TODO())
		})

		It("should apply the correct rate limit", func() {
			// Sleep one second
			time.Sleep(1 * time.Second)

			// Get config
			config, err := services.GetConfig(uid)
			Expect(err).To(BeNil())

			// Within rate / second
			for index := 0; index < config.Rate; index++ {
				status, err := services.Check(uid)
				Expect(err).To(BeNil())
				Expect(status).To(BeTrue())
			}

			// Above the rate / second limit
			status, err := services.Check(uid)
			Expect(err).To(BeNil())
			Expect(status).To(BeFalse())
		})

		Context("Delete", func() {
			It("should delete the config", func() {
				// Delete config
				services.DeleteConfig(uid)

				// Get config
				_, err := services.GetConfig(uid)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("Create", func() {
			It("should create the config", func() {
				// Create config
				services.CreateOrUpdateConfig(uid, *mockupFirstConfig)

				// Get config
				config, err := services.GetConfig(uid)
				Expect(err).To(BeNil())
				Expect(config).To(Equal(*mockupFirstConfig))
			})

			It("should update the config", func() {
				// Create config
				services.CreateOrUpdateConfig(uid, *mockupSecondConfig)

				// Get config
				config, err := services.GetConfig(uid)
				Expect(err).To(BeNil())
				Expect(config).To(Equal(*mockupSecondConfig))
			})
		})

		It("should apply the correct quota", func() {
			// Sleep one second
			time.Sleep(1 * time.Second)

			// Get config
			config, err := services.GetConfig(uid)
			Expect(err).To(BeNil())

			// Within quota / interval
			for index := 0; index < config.Quota.Number; index++ {
				status, err := services.Check(uid)
				Expect(err).To(BeNil())
				Expect(status).To(BeTrue())
			}

			// Above the quota / interval limit
			status, err := services.Check(uid)
			Expect(err).To(BeNil())
			Expect(status).To(BeFalse())
		})
	})
})
