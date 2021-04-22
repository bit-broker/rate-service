/*
 * File Name : log.go
 * Creation Date : 21-12-2020
 * Written by : Jean Diaconu <jdiaconu@cisco.com>
 * Copyright (C) Cisco System Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
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
