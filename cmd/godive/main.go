package main

import (
	"godive/pkg/entities/gradientfactor"
	"godive/pkg/entities/model"
	"godive/pkg/entities/tissue"
)

func main() {
	TissueConstructor := tissue.New
	gradientfactor := gradientfactor.New(0.35, 0.85)
	_ = model.NewZHL16C(TissueConstructor, gradientfactor)

}
