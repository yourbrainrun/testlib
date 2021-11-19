package main

import (
	"testing"
	"testlib/rand"
)

func TestRand(t *testing.T) {
	if rand.GetRand() != 0 {
	} else {
		t.Errorf("rand error")
	}
}
