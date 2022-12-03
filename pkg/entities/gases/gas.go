package gases

import (
	"fmt"
)

const (
	e = 1e-6
)

type Gas struct {
	ppO2       float32
	fO2        float32
	fHe        float32
	fN2        float32
	mod        float32
	startDepth float32
}

func NewGas(ppO2, pO2, pHe float32) (*Gas, error) {
	if pO2+pHe > 1 {
		return nil, fmt.Errorf("combined percententages exceeds 1 for pO2=%f and pHe%f", pO2, pHe)
	}
	g := &Gas{
		ppO2: ppO2,
		fO2:  pO2,
		fHe:  pHe,
	}

	g.CalculateN2()
	g.CalculateMOD()
	g.CalculateStartDepth()

	return g, nil
}

func (g *Gas) CalculateN2() {
	g.fN2 = 1.0 - (g.fO2 + g.fHe)
}

func (g *Gas) CalculateMOD() {
	g.mod = g.ppO2 / g.fO2
}

func (g *Gas) CalculateStartDepth() {
	if startDepth := 0.18 / g.fO2; startDepth > 1 {
		g.startDepth = startDepth
		return
	}
	g.startDepth = 1
}

func (g *Gas) MOD() float32 {
	return g.mod
}

func (g *Gas) StartDepth() float32 {
	return g.startDepth
}

func (g *Gas) getName() string {
	if floats.AlmostEqual(g.fO2, 0.21, e) && g.fHe == 0 {
		return "Air"
	}
	return "other"
}

func (g *Gas) String() string {
	name := g.getName()
	return name
}
