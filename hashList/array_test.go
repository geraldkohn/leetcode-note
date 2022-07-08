package main_test

import (
	"log"
	"testing"
)

func TestArray(t *testing.T) {
	m := make([][]int, 0)
	m = append(m, make([]int, 0))
	m[0] = append(m[0], 1)
	log.Println(m)
}