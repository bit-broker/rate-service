/*
 * File Name : models.go
 * Creation Date : 21-04-2021
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
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
