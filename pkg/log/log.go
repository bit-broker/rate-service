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
 * File Name : log.go
 * Creation Date : 21-12-2020
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 */

package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	formatter := &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}

	logrus.SetFormatter(formatter)
}

// SetLogLevel : Configure log level
func SetLogLevel(aLogLevel string) {
	// Set log level
	logLevel := logrus.InfoLevel
	if aLogLevel == "DebugLevel" {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
}

// Info : Configure log level
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Debug : Configure log level
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Fatal : Configure log level
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Error : Configure log level
func Error(args ...interface{}) {
	logrus.Error(args...)
}
