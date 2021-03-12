package main

import (
	"testing"

	"github.com/workmanager/parser"
)

func Test_methodExecution(t *testing.T) {
	parser.Reader("sanitaze.csv", parser.FieldNumber)

	parser.ReaderWithoutMetric("sanitaze.csv", parser.FieldNumber)
}
