package main

import (
	"fmt"
	"math/rand"
)

type MutationArg struct {
	minor     miniorMutationArg
	major     majorMutationArg
	majorRate float64
}

type miniorMutationArg struct {
	filp float64
	posi float64
	nega float64
}

type majorMutationArg struct {
	filp float64
	posi float64
	nega float64
}

func (cz *citizen) mutation() {
	sims := cz.BelongTo.simarg.seed
	rands := getPGSourceRand(sims + string(cz.BelongTo.Generation) + cz.getID() + string(cz.BelongTo.reproduceArg.creationSerial))
	multype := rands.Float64() + 0.5
	isminior := multype > cz.BelongTo.mutationArg.majorRate
	if isminior {
		cz.mutationFPNR(rands, cz.BelongTo.mutationArg.minor.filp, cz.BelongTo.mutationArg.minor.posi, cz.BelongTo.mutationArg.minor.nega)
	} else {
		cz.mutationFPNR(rands, cz.BelongTo.mutationArg.major.filp, cz.BelongTo.mutationArg.major.posi, cz.BelongTo.mutationArg.major.nega)
	}
	fmt.Printf("mutation %v %v %v\n", cz.getID(), isminior, cz.BelongTo.reproduceArg.creationSerial)
	cz.BelongTo.reproduceArg.creationSerial++
}

func (cz *citizen) mutationFPNR(rand *rand.Rand, filp float64, posi float64, nega float64) {
	var fcount, pcount, ncount, rcount int
	bnece := cz.BelongTo.Fitnesscalcarg.bph * cz.BelongTo.Fitnesscalcarg.bpw
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
