package driver

const (
	STATE_UNDEFINED uint = iota
	STATE_OFF
	STATE_ON
)

type Lamp struct {
	Pin      IPin
	tO       uint64
	tP       uint64
	tT       uint64
	cO       uint64
	cP       uint64
	cT       uint64
	reloaded bool
	state    uint
}

func NewLamp(pin IPin) *Lamp {
	l := &Lamp{Pin: pin, state: STATE_UNDEFINED}
	l.Pin.Output()
	l.turnOff()
	return l
}

//  ____    ____    __
// |    |__|    |__|   ...
// <-tO->
// <-tP--->
// <-tT------------------>
// tO - time of turn on
// tP - --//-- period
// tT - --//-- total
func (l *Lamp) Reload(tO, tP, tT uint64) {
	l.turnOff()
	l.tO = tO
	l.tP = tP
	l.tT = tT
	l.cO = tO
	l.cP = tP
	l.cT = tT
	l.reloaded = true
}

func (l *Lamp) Tick() {
	if l.reloaded { // first tick after reload - just setup initial state
		if l.cO > 0 {
			l.turnOn()
		} else {
			l.turnOff()
		}
		l.reloaded = false
		return
	}

	if l.cT > 0 { // if we have any total time limit
		l.cT--
		if l.cT == 0 {
			l.turnOff()
			l.cO, l.cP = 0, 0
			return
		}
	}

	if l.cO > 0 && l.cP > 0 { // both counting - first stage
		l.cO--
		l.cP--
		if l.cO == 0 { // end of ON state
			if l.cP == 0 { // if period is the same as ON time
				l.cO = l.tO
				l.cP = l.tP // just reload counters, state is the same
			} else { // period is greater - will be LO state
				l.turnOff() // don't reload counters - cP still counting
			}
		}
		return
	}

	if l.cP > 0 { // only cP counting
		l.cP--
		if l.cP == 0 {
			l.cO = l.tO
			l.cP = l.tP
			if l.cO > 0 {
				l.turnOn()
			}
		}
	}
}

func (l *Lamp) turnOn() {
	l.Pin.Low()
	l.state = STATE_ON
}

func (l *Lamp) turnOff() {
	l.Pin.High()
	l.state = STATE_OFF
}
