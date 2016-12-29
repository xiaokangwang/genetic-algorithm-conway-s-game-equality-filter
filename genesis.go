package main

import "math/big"

type genesisArg struct {
	density           float64
	initialPopulation int
	calcArg           fitnessCalcArg
}

const genesisPGentag = "genesis"

func (sc *society) Genesis(ga *genesisArg) {
	sc.Members = make([](*citizen), ga.initialPopulation, ga.initialPopulation)
	rand := getPGSourceRand(sc.simarg.seed + genesisPGentag)
	bnece := ga.calcArg.bph * ga.calcArg.bpw
	for cg := 0; cg < ga.initialPopulation; cg++ {
		cuCitizen := &citizen{Gene: big.NewInt(0)}
		for cub := 0; cub < bnece; cub++ {
			r := rand.Float64()
			if r > ga.density {
				cuCitizen.Gene.Add(cuCitizen.Gene, big.NewInt(1))
			}
			cuCitizen.Gene.Lsh(cuCitizen.Gene, 1)
		}
		cuCitizen.Gene.Rsh(cuCitizen.Gene, 1)
		cuCitizen.Fitness = notCalc
		cuCitizen.BelongTo = sc
		sc.Members[cg] = cuCitizen
	}
}
