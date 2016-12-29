package main

import "testing"

func TestGetPGSourceRand(t *testing.T) {
	s := getPGSourceRand("TestRand")
	var max, min float64
	for i := 1; i < 10000; i++ {
		c := s.NormFloat64() + 1
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
}
