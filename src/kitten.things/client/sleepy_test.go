package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSleepy(t *testing.T) {
	assert.Equal(t, 2, 2)
	time.Sleep(3 * time.Second)
}
