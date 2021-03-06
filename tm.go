package dyn4go

import (
	"math"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

func preambleCorrection(message string) string {
	_, f, l, r := runtime.Caller(2)
	if !r {
		return message
	}
	return "\r\t" + filepath.Base(f) + ":" + strconv.Itoa(l) + ": " + message
}

func AssertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Error(preambleCorrection("Values not equal in assertion"))
	}
}

func AssertEqualWithinError(t *testing.T, a, b, c float64) {
	if math.Abs(a)-math.Abs(b) > math.Abs(c) {
		t.Error(preambleCorrection("Values not equal within error in assertion"))
	}
}

func AssertTrue(t *testing.T, a bool) {
	if !a {
		t.Error(preambleCorrection("Condition not true in assertion"))
	}
}

func AssertNotEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Error(preambleCorrection("Values equal in assertion"))
	}
}

func AssertFalse(t *testing.T, a bool) {
	if a {
		t.Error(preambleCorrection("Condition not false in assertion"))
	}
}

/**
 * Usage:
 *
 * func TestSomething(t * testing.T) {
 *     defer AssertPanic(t)
 *     functionThatPanics(maybeSomeArguments)
 * }
 *
 */
func deferredPreambleCorrection(message string) string {
	_, f, l, r := runtime.Caller(2)
	if !r {
		return message
	}
	return "\r\t" + filepath.Base(f) + ":" + strconv.Itoa(l) + ": " + message
}

func AssertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Error(deferredPreambleCorrection("Function failed to panic"))
	}
}

func AssertNoPanic(t *testing.T) {
	if r := recover(); r != nil {
		t.Error(deferredPreambleCorrection("Function paniced"))
	}
}
