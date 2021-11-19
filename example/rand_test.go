package main

import (
	"github.com/yourbrainrun/testlib/my_rand"
	"testing"
)

func TestRand(t *testing.T) {
	if my_rand.GetRand() != 0 {
	} else {
		t.Errorf("my_rand error")
	}
}
