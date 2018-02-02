package main

import (
	"github.com/scele/go-dep-bazel-vscode-example/pkg/something"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info(something.GetMessage())
}
