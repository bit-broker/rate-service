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
 * File Name : error.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package helper

import (
	"encoding/json"
	"net/http"

	"github.com/bit-broker/rate-service/pkg/log"
)

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare internal standard error.
func GetError(err error, w http.ResponseWriter) {
	log.Error(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	_, _ = w.Write(message)
}

// GetNotFoundError : This is helper function to prepare not found error.
func GetNotFoundError(w http.ResponseWriter) {
	var response = ErrorResponse{
		ErrorMessage: "Not Found",
		StatusCode:   http.StatusNotFound,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	_, _ = w.Write(message)
}

// GetBadRequestError : This is helper function to prepare bad request error.
func GetBadRequestError(w http.ResponseWriter) {
	var response = ErrorResponse{
		ErrorMessage: "Bad request",
		StatusCode:   http.StatusBadRequest,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	_, _ = w.Write(message)
}
