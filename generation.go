package main

import "fmt"

func (sc *society) generation() {
	sc.calcAdjustmentFactor()
	newMember := make([](*citizen), 0, int(sc.RreproduceArg.AdjustmentFactor*float64(len(sc.Members))))
	for rank, citizennow := range sc.Members {
		rpc := citizennow.ReproduceCount(rank + 1)
		//fmt.Printf("reps %v\n", rpc)
		for cr := 0; cr < rpc; cr++ {
			cz := citizennow.ReproduceOnce()
			fmt.Printf("parent %v -> %v\n", cz.getID(), citizennow.getID())
			newMember = append(newMember, cz)
		}
	}
	sc.Generation++
	sc.Members = newMember
}

func (sc *society) legacy() {
	sc.sortCitizen()
	sc.Achievement = sc.Members[0].getTrace()
}
