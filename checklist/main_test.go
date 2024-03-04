package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	data := []struct {
		a, b, want float64
	}{
		{
			a:    5.0,
			b:    5.5,
			want: 10.5,
		},
		{
			a:    2.2,
			b:    2.0,
			want: 4.2,
		},
	}

	for _, d := range data {
		got := sum(d.a, d.b)
		want := d.want
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	}
}

func TestAddition(t *testing.T) {
	adds := []struct {
		value1, value2, want float64
	}{
		{
			value1: 2.2,
			value2: 4.0,
			want:   6.2,
		},
		{
			value1: 4.2,
			value2: 9.0,
			want:   13.2,
		},
	}
	for _, d := range adds {
		add := Addition{value1: d.value1, value2: d.value2}
		got := add.addition()
		want := d.want
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	}
}

func TestSubtraction(t *testing.T) {
	subs := []struct {
		value1, value2, want int
	}{
		{
			value1: 2,
			value2: 1,
			want:   1,
		},
		{
			value1: 10,
			value2: 8,
			want:   2,
		},
	}
	for _, d := range subs {
		got, err := subtract(d.value1, d.value2)
		if err != nil {
			t.Errorf("Error %v", err)
		}
		want := d.want
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestHandleSwitchBlock(t *testing.T) {
	data := []struct {
		value string
		want  int
	}{
		{
			value: "one",
			want:  1,
		},
		{
			value: "two",
			want:  2,
		},
		{
			value: "three",
			want:  3,
		},
		{
			value: "four",
			want:  0,
		},
	}

	for _, d := range data {
		got := handleSwitchBlock(d.value)
		want := d.want
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestHandleTime(t *testing.T) {
	data := []struct {
		dateS, want string
	}{
		{
			dateS: "22-02-2024",
			want:  "February",
		},
		{
			dateS: "22-01-2024",
			want:  "January",
		},
		{
			dateS: "12-12-2023",
			want:  "December",
		},
	}
	for _, d := range data {
		got, err := FetchMonthName(d.dateS)
		if err != nil {
			t.Errorf("Error %v", err)
		}
		want := d.want
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
