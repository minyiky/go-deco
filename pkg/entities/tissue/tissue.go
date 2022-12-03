package tissue

type Constructor func(
	aN2, bN2, aHe, bHe float32,
	n2HalfLife, heHalfLife float32,
) Tissue

type Tissue interface {
	Load()
	CalculateCeiling(float32, float32, float32) float32
}

type BuhlmannTissue struct {
	aN2, bN2, aHe, bHe     float32
	N2HalfLife, HeHalfLife float32
}

var New Constructor = func(
	aN2, bN2, aHe, bHe float32,
	n2HalfLife, heHalfLife float32,
) Tissue {
	return &BuhlmannTissue{
		aN2:        aN2,
		aHe:        aHe,
		bN2:        bN2,
		bHe:        bHe,
		N2HalfLife: n2HalfLife,
		HeHalfLife: heHalfLife,
	}
}

func (t *BuhlmannTissue) Load() {}

func (t *BuhlmannTissue) CalculateCeiling(gf, pN2, pHe float32) float32 {
	p := pN2 + pHe
	a := (t.aN2*pN2 + t.aHe*pHe) / p
	b := (t.bN2*pN2 + t.bHe*pHe) / p
	return (p - a*gf) / (gf/b + 1.0 - gf)
}
