package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(10, 5)
	want := Number(15)

	if want != got {
		t.Errorf("got %d want %d given", got, want)
	}
}
