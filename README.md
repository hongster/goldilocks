# Goldilocks

Simple wrapper for Go's [log.Logger](https://golang.org/pkg/log/#Logger) that logs to Stdout and Stderr by default, and is customizable. By default informational and warning message goes to Stdout, and error message goes to Stderr.

Example

	package main

	import (
		"github.com/hongster/goldilocks"
	)

	func main() {
		logger := goldilocks.NewLogger()

		logger.Info("Testing")
		logger.Infof("%v Bar", "Foo")

		logger.Warning("Testing")
		logger.Warningf("%v Bar", "Foo")

		logger.Error("Testing")
		logger.Errorf("%v Bar", "Foo")
	}

Resultant output

	INFO: 2016/04/26 01:02:02 Testing
	INFO: 2016/04/26 01:02:02 Foo Bar
	WARNING: 2016/04/26 01:02:02 Testing
	WARNING: 2016/04/26 01:02:02 Foo Bar
	ERROR: 2016/04/26 01:02:02 Testing
	ERROR: 2016/04/26 01:02:02 Foo Bar
