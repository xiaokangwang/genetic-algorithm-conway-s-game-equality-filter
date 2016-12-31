package main

import (
	"fmt"
	"math"
	"math/big"
)

type ReproduceArg struct {
	AdjustmentFactor float64
	CreationSerial   int
	LuckyFactor      *float64
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func (ctz *citizen) ReproduceCount(rank int) int {
	count := len(ctz.belongTo.Members)
	probirep := calcScore(ctz.belongTo.Equality, float64(count), float64(rank)) * ctz.belongTo.RreproduceArg.AdjustmentFactor
	sims := ctz.belongTo.Ssimarg.Seed
	rands := getPGSourceRand(sims + string(ctz.belongTo.Generation) + ctz.getID() + "rep")
	r := math.Abs(rands.NormFloat64()**ctz.belongTo.RreproduceArg.LuckyFactor) + 1
	co := probirep / r
	fmt.Printf("rep %v / %v %v %v %v %v\n", probirep, r, co, rank, ctz.getID(), ctz.getFitness())
	return int(round(co))
}

func calcScore(Equality, count, rank float64) float64 {
	n := 1 / Equality
	probirep := math.Exp((n)*math.Log(((count)/(rank)))) / (((-count) + math.Exp(math.Log((count)*(n)))) / (-1 + (n)))
	return probirep
}

func (ctz *citizen) ReproduceOnce(firstChild bool, serial int) *citizen {
	cuCtz := &citizen{}
	*cuCtz = *ctz

	cuCtz.Gene = big.NewInt(0)
	cuCtz.Gene.Set(ctz.Gene)

	/*
		fmt.Println(debugTraceStart)
		fmt.Println("mutation Before")
		spew.Dump(cuCtz.getID())
		fmt.Println(debugTraceEnd)
	*/

	if !firstChild {
		cuCtz.mutation(serial)
	}

	/*
		fmt.Println(debugTraceStart)
		fmt.Println("mutation After")
		spew.Dump(cuCtz.getID())
		fmt.Println(debugTraceEnd)
	*/
	return cuCtz
}
