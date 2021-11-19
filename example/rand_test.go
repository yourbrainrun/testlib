package main

import (
	"testing"
	"testlib/my_rand"
)

func TestRand(t *testing.T) {
	if my_rand.GetRand() != 0 {
	} else {
		t.Errorf("my_rand error")
	}
}
