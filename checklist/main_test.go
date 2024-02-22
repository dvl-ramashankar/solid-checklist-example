package main

import "testing"

func TestSum(t *testing.T) {
	data := []map[string]interface{}{
		{
			"a":    5.0,
			"b":    5.5,
			"want": 10.5,
		},
		{
			"a":    2.2,
			"b":    2.0,
			"want": 4.2,
		},
	}

	for _, d := range data {
		got := sum(d["a"].(float64), d["b"].(float64))
		want := d["want"].(float64)
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	}
}

func TestAddition(t *testing.T) {
	adds := []map[string]interface{}{
		{
			"value1": 2.2,
			"value2": 4.0,
			"want":   6.2,
		},
		{
			"value1": 4.2,
			"value2": 9.0,
			"want":   13.2,
		},
	}
	for _, d := range adds {
		add := Addition{value1: d["value1"].(float64), value2: d["value2"].(float64)}
		got := add.addition()
		want := d["want"].(float64)
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	}
}
