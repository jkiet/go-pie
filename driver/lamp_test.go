package driver

import (
	"testing"
)

func TestInit(t *testing.T) {
	lamp := NewLamp(NewPin())
	if lamp.state != STATE_OFF {
		t.Errorf("Wrong state after pin initialization: %v", lamp.state)
	}
}

func TestTurnOnForever(t *testing.T) {
	lamp := NewLamp(NewPin())
	lamp.Reload(1, 1, 0)
	if lamp.state != STATE_OFF {
		t.Errorf("Wrong state after pin initialization: %v", lamp.state)
	}
	lamp.Tick()
	if lamp.state != STATE_ON {
		t.Errorf("Wrong state after initial tick: %v", lamp.state)
	}
	for i := 1; i < 3; i++ {
		lamp.Tick()
		if lamp.state != STATE_ON {
			t.Errorf("Wrong state after %v tick: %v", i, lamp.state)
		}
	}
}

func TestTurnOff(t *testing.T) {
	lamp := NewLamp(NewPin())
	lamp.Reload(1, 1, 0)
	if lamp.state != STATE_OFF {
		t.Errorf("Wrong state after pin initialization: %v", lamp.state)
	}
	lamp.Tick()
	if lamp.state != STATE_ON {
		t.Errorf("Wrong state after initial tick: %v", lamp.state)
	}
	lamp.Reload(0, 0, 0)
	for i := 1; i < 3; i++ {
		lamp.Tick()
		if lamp.state != STATE_OFF {
			t.Errorf("Wrong state after %v tick: %v", i, lamp.state)
		}
	}
}

func TestTurnOnLimited(t *testing.T) {
	lamp := NewLamp(NewPin())
	lamp.Reload(1, 1, 5)
	if lamp.state != STATE_OFF {
		t.Errorf("Wrong state after pin initialization: %v", lamp.state)
	}
	lamp.Tick()
	if lamp.state != STATE_ON {
		t.Errorf("Wrong state after initial tick: %v", lamp.state)
	}
	for i := 1; i < 5; i++ {
		lamp.Tick()
		if lamp.state != STATE_ON {
			t.Errorf("Wrong state after %v tick: %v", i, lamp.state)
		}
	}
	for i := 5; i < 10; i++ {
		lamp.Tick()
		if lamp.state != STATE_OFF {
			t.Errorf("Wrong state after %v tick: %v", i, lamp.state)
		}
	}
}

func TestBlinkForever(t *testing.T) {
	lamp := NewLamp(NewPin())
	lamp.Reload(1, 2, 0)
	if lamp.state != STATE_OFF {
		t.Errorf("Wrong state after pin initialization: %v", lamp.state)
	}
	for i := 0; i < 10; i++ {
		lamp.Tick()
		if lamp.state != STATE_ON {
			t.Errorf("Wrong state after %v tick: %v", i*2, lamp.state)
		}
		lamp.Tick()
		if lamp.state != STATE_OFF {
			t.Errorf("Wrong state after %v tick: %v", (i*2)+1, lamp.state)
		}
	}
}

func TestBlinkLimited(t *testing.T) {
	lamp := NewLamp(NewPin())
	lamp.Reload(1, 2, 10)
	if lamp.state != STATE_OFF {
		t.Errorf("Wrong state after pin initialization: %v", lamp.state)
	}
	for i := 0; i < 5; i++ {
		lamp.Tick()
		if lamp.state != STATE_ON {
			t.Errorf("Wrong state after %v tick: %v", i*2, lamp.state)
		}
		lamp.Tick()
		if lamp.state != STATE_OFF {
			t.Errorf("Wrong state after %v tick: %v", (i*2)+1, lamp.state)
		}
	}
	for i := 10; i < 15; i++ {
		lamp.Tick()
		if lamp.state != STATE_OFF {
			t.Errorf("Wrong state after %v tick: %v", i, lamp.state)
		}
	}
}
