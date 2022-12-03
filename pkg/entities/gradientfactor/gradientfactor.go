package gradientfactor

type GradientFactor struct {
	low, high float32
}

func New(low, high float32) *GradientFactor {
	return &GradientFactor{
		low:  low,
		high: high,
	}
}

func (gf *GradientFactor) Low() float32 {
	return gf.low
}

func (gf *GradientFactor) High() float32 {
	return gf.low
}

func (gf *GradientFactor) SetLow(val float32) {
	gf.low = val
}

func (gf *GradientFactor) SetHigh(val float32) {
	gf.high = val
}

func (gf *GradientFactor) Set(low, high float32) {
	gf.SetLow(low)
	gf.SetHigh(high)
}
