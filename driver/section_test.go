package driver

import (
	"testing"
)

func TestParseDefaults(t *testing.T) {
	section := NewSection()
	s, l, tO, tP, tT, err := section.parseKV("2", "1")
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	if s != SECTION {
		t.Errorf("Expected section: %v ; got: %v", SECTION, s)
		return
	}
	if l != 2 {
		t.Errorf("Expected lamp: 2 ; got: %v", l)
		return
	}
	if tO != 1 {
		t.Errorf("Expected tH: 1 ; got: %v", tO)
		return
	}
	if tP != tO {
		t.Errorf("Expected tP: %v ; got: %v", tO, tP)
		return
	}
	if tT != 0 {
		t.Errorf("Expected tT: 0 ; got: %v", tT)
		return
	}
}

func TestParseSection(t *testing.T) {
	section := NewSection()
	s, _, _, _, _, err := section.parseKV("7-2", "1")
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	if s != 7 {
		t.Errorf("Expected section: 7 ; got: %v", s)
		return
	}
}

func TestParseTotalTime(t *testing.T) {
	section := NewSection()
	_, _, _, _, tT, err := section.parseKV("0-2", "1 500")
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	if tT != 500 {
		t.Errorf("Expected tT: 500 ; got: %v", tT)
		return
	}
}

func TestParsePeriodTime(t *testing.T) {
	section := NewSection()
	_, _, _, tP, _, err := section.parseKV("0-2", "100/200 1000")
	if err != nil {
		t.Errorf("Got error: %v", err)
		return
	}
	if tP != 200 {
		t.Errorf("Expected tP: 200 ; got: %v", tP)
		return
	}
}
