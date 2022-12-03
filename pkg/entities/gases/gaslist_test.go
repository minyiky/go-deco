package gases_test

import (
	"fmt"
	"godive/pkg/entities/gases"
	"sort"
	"testing"
)

func TestXxx(t *testing.T) {
	g1, _ := gases.NewGas(1.4, 0.21, 0)
	g2, _ := gases.NewGas(1.6, 0.80, 0)

	gl := gases.NewGasList()
	gl.Add(*g2)
	gl.Add(*g1)

	fmt.Println(gl)

	sort.Sort(gl)

	fmt.Println(gl)

	t.FailNow()
}
