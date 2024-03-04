package main

import "testing"

func TestVec2Magnitude(t *testing.T) {
	v := New(4, 0)
	got := v.Magnitude()
	want := 4.0
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestVec2MagnitudeBothAxes(t *testing.T) {
	v := New(4, -3)
	got := v.Magnitude()
	want := 5.0
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
