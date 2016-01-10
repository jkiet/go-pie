package driver

import (
	"errors"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"strconv"
	"strings"
	"time"
)

const (
	SECTION uint64 = 0
)

// our order to PI GPIO
var layout = map[uint64]rpio.Pin{
	0: 2, // GPIO2
	1: 3,
	2: 4,
}

type Section struct {
	Lamps map[uint64]*Lamp
}

func NewSection() *Section {
	return &Section{}
}

func (s *Section) Init() (err error) {
	err = rpio.Open()
	if err != nil {
		return
	}
	s.Lamps = make(map[uint64]*Lamp)
	for l, p := range layout {
		s.Lamps[l] = NewLamp(rpio.Pin(p))
	}
	ticker := time.NewTicker(time.Millisecond * 100).C
	go func() {
		for {
			select {
			case <-ticker:
				for _, l := range s.Lamps {
					l.Tick()
				}
			}
		}
	}()
	return
}

func (this *Section) Reload(commands map[string]string) (status map[string]string) {
	status = make(map[string]string)
	for k, v := range commands {
		s, l, tO, tP, tT, err := this.parseKV(k, v)
		if err != nil {
			status[k] = fmt.Sprintf("Parse error (%v=>%v) : %v", k, v, err)
			continue
		}
		if s != SECTION {
			status[k] = fmt.Sprintf("I'm not target section (this section: %v , target section: %v)", SECTION, s)
			continue
		}
		if lamp, exists := this.Lamps[l]; exists {
			lamp.Reload(tO, tP, tT)
			status[k] = "OK"
		} else {
			status[k] = fmt.Sprintf("Unable to find lamp %v)", l)
		}
	}
	return
}

func (this *Section) parseKV(k, v string) (s, l, tO, tP, tT uint64, err error) {
	// parse key
	if strings.Contains(k, "-") {
		kArr := strings.Split(k, "-")
		if len(kArr) != 2 {
			err = errors.New("can't parse key " + k)
			return
		}
		s, err = strconv.ParseUint(kArr[0], 10, 32)
		if err != nil {
			return
		}
		l, err = strconv.ParseUint(kArr[1], 10, 32)
		if err != nil {
			return
		}
	} else {
		s = SECTION
		l, err = strconv.ParseUint(k, 10, 32)
		if err != nil {
			return
		}
	}

	// parse value
	tOtP := v
	tT = 0
	if strings.Contains(v, " ") {
		vArr := strings.Split(v, " ")
		if len(vArr) != 2 {
			err = errors.New("can't parse value " + v)
			return
		}
		tOtP = vArr[0]
		tT, err = strconv.ParseUint(vArr[1], 10, 32)
		if err != nil {
			return
		}
	}

	if strings.Contains(tOtP, "/") {
		tArr := strings.Split(tOtP, "/")
		if len(tArr) != 2 {
			err = errors.New("can't parse value " + v)
			return
		}
		tO, err = strconv.ParseUint(tArr[0], 10, 32)
		if err != nil {
			return
		}
		tP, err = strconv.ParseUint(tArr[1], 10, 32)
		if err != nil {
			return
		}
	} else {
		tO, err = strconv.ParseUint(tOtP, 10, 32)
		if err != nil {
			return
		}
		tP = tO
	}
	return
}
