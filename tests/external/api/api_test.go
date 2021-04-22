/*
 * File Name : api_test.go
 * Creation Date : 05-05-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

package tests

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/bit-broker/rate-service/internal/helper"
	"github.com/bit-broker/rate-service/internal/models"
	"github.com/bit-broker/rate-service/internal/routes"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// ------------------------ GLOBAL -------------------- //

var uid = strconv.Itoa(rand.Intn(100))
var mockupFirstConfig = `{"enabled":true,"quota":{"max_number":10,"interval_type":"month"},"rate":1}`
var mockupSecondConfig = `{"enabled":true,"quota":{"max_number":15,"interval_type":"day"},"rate":2}`

// ------------------------ GLOBAL -------------------- //

// TestAPI : API Test cases
func TestAPI(t *testing.T) {
	// Load env
	helper.LoadEnv(helper.TestEnv)

	RegisterFailHandler(Fail)
	RunSpecs(t, "API Test Suite")
}

var _ = Describe("API", func() {
	var (
		router *mux.Router
	)

	BeforeEach(func() {
		router = routes.InitializeRouter()
	})

	Context("Health Check", func() {
		It("should be up", func() {
			// Create request
			req, err := http.NewRequest("GET", "/api/v1", nil)
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))
		})
	})

	Context("Rate Config Routes", func() {
		It("shouldn't find the config", func() {
			// Create request
			req, err := http.NewRequest("GET", "/api/v1/"+uid+"/config", nil)
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusNotFound))
		})

		It("should set a new config", func() {
			// Create request
			var jsonData = []byte(mockupFirstConfig)
			req, err := http.NewRequest("PUT", "/api/v1/"+uid+"/config", bytes.NewBuffer(jsonData))
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))
		})

		It("should find the right config", func() {
			// Create request
			req, err := http.NewRequest("GET", "/api/v1/"+uid+"/config", nil)
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))

			// Check config
			var config, initialConfig models.Config
			err = json.NewDecoder(rr.Body).Decode(&config)
			json.Unmarshal([]byte(mockupFirstConfig), &initialConfig)
			Expect(err).To(BeNil())
			Expect(config).To(Equal(initialConfig))

		})

		It("should override with a new config", func() {
			// Create request
			var jsonData = []byte(mockupSecondConfig)
			req, err := http.NewRequest("PUT", "/api/v1/"+uid+"/config", bytes.NewBuffer(jsonData))
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))
		})

		It("should find the right config", func() {
			// Create request
			req, err := http.NewRequest("GET", "/api/v1/"+uid+"/config", nil)
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))

			// Check config
			var config, initialConfig models.Config
			err = json.NewDecoder(rr.Body).Decode(&config)
			json.Unmarshal([]byte(mockupSecondConfig), &initialConfig)
			Expect(err).To(BeNil())
			Expect(config).To(Equal(initialConfig))

		})

		It("should delete the config", func() {
			// Create request
			req, err := http.NewRequest("DELETE", "/api/v1/"+uid+"/config", nil)
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))
		})

		It("should find an empty config", func() {
			// Create request
			req, err := http.NewRequest("GET", "/api/v1/"+uid+"/config", nil)
			Expect(err).To(BeNil())

			// Create recorder
			rr := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusNotFound))
		})
	})
})
