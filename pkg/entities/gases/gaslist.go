package gases

type GasList struct {
	gases []Gas
}

func NewGasList() *GasList {
	return &GasList{}
}

func (gl *GasList) Add(g Gas) {
	gl.gases = append(gl.gases, g)
}

func (gl *GasList) Len() int {
	return len(gl.gases)
}

func (gl *GasList) Less(i, j int) bool {
	return gl.gases[i].mod > gl.gases[j].mod
}

func (gl *GasList) Swap(i, j int) {
	gl.gases[i], gl.gases[j] = gl.gases[j], gl.gases[i]
}
