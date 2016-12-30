package main

import (
	"fmt"
	"math/rand"
)

type MutationArg struct {
	Mminor     MiniorMutationArg
	Mmajor     MajorMutationArg
	MmajorRate float64
}

type MiniorMutationArg struct {
	Filp float64
	Posi float64
	Nega float64
}

type MajorMutationArg struct {
	Filp float64
	Posi float64
	Nega float64
}

func (cz *citizen) mutation(serial int) {
	sims := cz.belongTo.Ssimarg.Seed
	rands := getPGSourceRand(sims + string(cz.belongTo.Generation) + cz.getID() + string(serial))
	multype := rands.Float64() + 0.5
	isminior := multype > cz.belongTo.MmutationArg.MmajorRate
	if isminior {
		cz.mutationFPNR(rands, cz.belongTo.MmutationArg.Mminor.Filp, cz.belongTo.MmutationArg.Mminor.Posi, cz.belongTo.MmutationArg.Mminor.Nega)
	} else {
		cz.mutationFPNR(rands, cz.belongTo.MmutationArg.Mmajor.Filp, cz.belongTo.MmutationArg.Mmajor.Posi, cz.belongTo.MmutationArg.Mmajor.Nega)
	}
	fmt.Printf("mutationO %v %v %v\n", cz.getID(), isminior, serial)
}

func (cz *citizen) mutationFPNR(rand *rand.Rand, filp float64, posi float64, nega float64) {
	var fcount, pcount, ncount, rcount int
	bnece := cz.belongTo.Fitnesscalcarg.Bph * cz.belongTo.Fitnesscalcarg.Bpw
	for cub := 0; cub < bnece; cub++ {
		r := rand.Float64() + 0.5
		if r < filp {
			b := cz.Gene.Bit(cub)
			if b == 0 {
				b = 1
			} else {
				b = 0
			}
			cz.Gene = cz.Gene.SetBit(cz.Gene, cub, b)
			cz.Fitness = notCalc
			fcount++
			continue
		}
		if r < posi+filp {
			cz.Gene = cz.Gene.SetBit(cz.Gene, cub, 1)
			cz.Fitness = notCalc
			pcount++
			continue
		}
		if r < filp+posi+nega {
			cz.Gene = cz.Gene.SetBit(cz.Gene, cub, 0)
			cz.Fitness = notCalc
			ncount++
			continue
		}
		rcount++
	}
	fmt.Printf("mutationFPNR %v %v %v %v %v\n", cz.getID(), fcount, pcount, ncount, rcount)
}
