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
 * File Name : models.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package models

// IntervalType : Type of interval
type IntervalType string

// Day
// Month
const (
	DayType   IntervalType = "day"
	MonthType IntervalType = "month"
)

// Quota Struct
type Quota struct {
	Number   int          `json:"max_number,omitempty" bson:"max_number,omitempty"`
	Interval IntervalType `json:"interval_type,omitempty" bson:"interval_type,omitempty"`
}

// Config Struct
type Config struct {
	Enabled bool           `json:"enabled" bson:"enabled"`
	Quota   Quota          `json:"quota,omitempty" bson:"quota,omitempty"`
	Rate    int            `json:"rate,omitempty" bson:"rate,omitempty"`
	Log     map[string]int `json:"log,omitempty" bson:"log,omitempty"`
}
