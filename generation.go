package main

import (
	"fmt"
	"sync"
)

func (sc *society) generation() {
	sc.calcAdjustmentFactor()
	newMember := make([](*citizen), 0, int(sc.RreproduceArg.AdjustmentFactor*float64(len(sc.Members))))
	sy := &sync.WaitGroup{}
	memberlock := &sync.Mutex{}
	memberlock.Lock()
	var count int
	for rank, citizennow := range sc.Members {
		rpc := citizennow.ReproduceCount(rank + 1)
		//fmt.Printf("reps %v\n", rpc)
		for cr := 0; cr < rpc; cr++ {
			sy.Add(1)
			fc := cr == 0
			go func(isFirstChildCurrent bool, CreationSeriaL int, syncs *sync.WaitGroup, czn *citizen, ml *sync.Mutex, index int) {
				cz := czn.ReproduceOnce(isFirstChildCurrent, CreationSeriaL)
				fmt.Printf("parent %v -> %v\n", cz.getID(), czn.getID())
				/*
					We are intented to make all result from our experiment stable and reproduceable
					By sorting result by genid before sorting with Fitness,
					We have elimated the uncertainty introduced by append result to newMember array
					without order(which= come with race).

					The compution cost in sorting is barely noticeable.
				*/
				cz.genid = index
				ml.Lock()
				newMember = append(newMember, cz)
				ml.Unlock()
				sy.Done()
			}(fc, sc.RreproduceArg.CreationSerial, sy, citizennow, memberlock, count)
			sc.RreproduceArg.CreationSerial++
		}

	}
	memberlock.Unlock()
	sy.Wait()
	sc.Generation++
	sc.Members = newMember
}

func (sc *society) legacy() {
	sc.sortCitizen()
	sc.Achievement = sc.Members[0].getTrace()
}
