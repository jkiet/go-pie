package config

import "testing"

func TestConfig(t *testing.T) {
	var cfg = `
section: 5
layout: [2, 4, 64]
`
	c := NewConfig()
	err := c.Read([]byte(cfg))
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	if c.Section != 5 {
		t.Errorf("Expected section value: 5 ; got: %v", c.Section)
		return
	}
	if len(c.Layout) != 3 {
		t.Errorf("Expected 3 elements in array; got: %v", len(c.Layout))
		return
	}
	if c.Layout[0] != 2 || c.Layout[1] != 4 || c.Layout[2] != 64 {
		t.Errorf("Expected array elements are: 2, 4, 64 ; got: %v, %v, %v", c.Layout[0], c.Layout[1], c.Layout[2])
		return
	}
}
