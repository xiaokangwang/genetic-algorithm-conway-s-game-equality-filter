package main

import "math/big"

type genesisArg struct {
	Density           float64
	InitialPopulation int
}

const genesisPGentag = "genesis"

func (sc *society) Genesis(ga *genesisArg) {
	sc.Members = make([](*citizen), ga.InitialPopulation, ga.InitialPopulation)
	rand := getPGSourceRand(sc.Ssimarg.Seed + genesisPGentag)
	bnece := sc.Fitnesscalcarg.Bph * sc.Fitnesscalcarg.Bpw
	for cg := 0; cg < ga.InitialPopulation; cg++ {
		cuCitizen := &citizen{Gene: big.NewInt(0)}
		for cub := 0; cub < bnece; cub++ {
			r := rand.Float64()
			if r > ga.Density {
				cuCitizen.Gene.Add(cuCitizen.Gene, big.NewInt(1))
			}
			cuCitizen.Gene.Lsh(cuCitizen.Gene, 1)
		}
		cuCitizen.Gene.Rsh(cuCitizen.Gene, 1)
		cuCitizen.Fitness = notCalc
		cuCitizen.belongTo = sc
		sc.Members[cg] = cuCitizen
	}
}
