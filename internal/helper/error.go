/*
 * File Name : error.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
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
