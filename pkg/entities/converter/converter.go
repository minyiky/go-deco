package converter

type Converter struct{
	waterDensity float32
}

func New(
	waterDensity float32,
) *Converter{
	return &Converter{
		waterDensity: waterDensity,
	}
}

func (c *Converter) DepthToPressure(depth float32) float32{
	return (depth * c.waterDensity / 10) + 1
}

func (c *Converter) PressureToDepth(pressure float32) float32{
	return (pressure - 1) * 10 / c.waterDensity
}
