package api

import (
	"testing"
)

func NoTestFunc() string {
	return "test"
}

func TestGetMethodName(t *testing.T) {
	m := getMethodName(NoTestFunc)
	t.Log(m)
}
