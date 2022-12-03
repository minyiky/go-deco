package step

import (
	"fmt"
	"time"
)

type phase string

const (
	Ascending phase = "Ascending"
	Decending phase = "Decending"
)

type Gas interface {
	Name() string
}

type Data interface {
	GF() float32
}

type Step struct {
	phase phase
	absP  float32
	time  time.Duration
	gas   Gas
	data  Data
}

func New(
	phase phase,
	absP float32,
	time time.Duration,
	gas Gas,
	data Data,
) *Step {
	return &Step{
		phase: phase,
		absP:  absP,
		time:  time,
		gas:   gas,
		data:  data,
	}
}

func (s *Step) String() string {
	return fmt.Sprintf("Step(phase=%s, abs_p=%.4f, time=%.4f, gas=%s, gf=%.2f)", s.phase, s.absP, s.time.Minutes(), s.gas.Name(), s.data.GF())
}
