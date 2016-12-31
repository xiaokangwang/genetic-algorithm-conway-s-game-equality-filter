package main

import (
	"fmt"
	"math/big"
	"sort"
	"sync"

	"golang.org/x/crypto/sha3"
)

type society struct {
	Members        [](*citizen)
	Achievement    int
	Generation     int
	Equality       float64
	Genesisarg     *genesisArg
	Fitnesscalcarg *fitnessCalcArg
	Ssimarg        *SimGlobalArg
	MmutationArg   *MutationArg
	RreproduceArg  *ReproduceArg
}

type citizen struct {
	belongTo *society
	Gene     *big.Int
	Fitness  int
	genid    int
}

const fitnesso = "fitness %v %v\n"

func (ctz *citizen) getFitness() int {
	if ctz.Fitness == notCalc {
		ctz.Fitness = getFitness(ctz.Gene, ctz.belongTo.Fitnesscalcarg)
	}
	fmt.Printf(fitnesso, ctz.getID(), ctz.Fitness)
	return ctz.Fitness
}

const tr = "checkpoint/%v/%v-%v"

func (ctz *citizen) getTrace() int {
	filename := fmt.Sprintf(tr, ctz.belongTo.Ssimarg.Seed, ctz.belongTo.Generation, ctz.getID())
	ctz.Fitness = getFitnessTrace(ctz.Gene, ctz.belongTo.Fitnesscalcarg, filename)

	return ctz.Fitness
}

func (ctz *citizen) getID() string {
	var id [16]byte
	sha3.ShakeSum128(id[:], ctz.Gene.Bytes())
	return fmt.Sprintf("%x", id)
}

func (sc *society) outputMember() {
	for _, cuctz := range sc.Members {
		fmt.Printf("socin %v -> %v\n", cuctz.getID(), sc.Generation)
	}
}

func (sc *society) outputmeta() {
	s := sc
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.bph", s.Fitnesscalcarg.Bph)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.bpw", s.Fitnesscalcarg.Bpw)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.height", s.Fitnesscalcarg.Height)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.width", s.Fitnesscalcarg.Width)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.step", s.Fitnesscalcarg.Step)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.offseth", s.Fitnesscalcarg.Offseth)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.offsetw", s.Fitnesscalcarg.Offsetw)
	fmt.Printf("Metadata %v %v\n", "Genesisarg.initialPopulation", s.Genesisarg.InitialPopulation)
	fmt.Printf("Metadata %v %v\n", "Genesisarg.density", s.Genesisarg.Density)
	fmt.Printf("Metadata %v %v\n", "mutationArg.major.filp", s.MmutationArg.Mmajor.Filp)
	fmt.Printf("Metadata %v %v\n", "mutationArg.major.nega", s.MmutationArg.Mmajor.Nega)
	fmt.Printf("Metadata %v %v\n", "mutationArg.major.posi", s.MmutationArg.Mmajor.Posi)
	fmt.Printf("Metadata %v %v\n", "mutationArg.minor.filp", s.MmutationArg.Mminor.Filp)
	fmt.Printf("Metadata %v %v\n", "mutationArg.minor.nega", s.MmutationArg.Mminor.Nega)
	fmt.Printf("Metadata %v %v\n", "mutationArg.minor.posi", s.MmutationArg.Mminor.Posi)
	fmt.Printf("Metadata %v %v\n", "mutationArg.majorRate", s.MmutationArg.MmajorRate)
	fmt.Printf("Metadata %v %v\n", "ReproduceArg.LuckyFactor", s.RreproduceArg.LuckyFactor)
	fmt.Printf("Metadata %v %v\n", "simarg.seed", s.Ssimarg.Seed)
	fmt.Printf("Metadata %v %v\n", "Equality", s.Equality)
	fmt.Printf("Metadata %v %v\n", "Ver", 4)
}

func (sc *society) calcMemberFitness() {
	sy := &sync.WaitGroup{}
	for _, mem := range sc.Members {
		sy.Add(1)
		go func(ctzcalc *citizen, sy *sync.WaitGroup) {
			ctzcalc.getFitness()
			sy.Done()
		}(mem, sy)
	}
	sy.Wait()
}

func (sc *society) sortCitizen() {
	//calcFitness
	sc.calcMemberFitness()
	spct := &societySorterGenid{toSort: sc}
	sort.Sort(sort.Reverse(spct))
	sct := &societySorterByFitness{toSort: sc}
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Sort Before")
		spew.Dump(sct)
		fmt.Println(debugTraceEnd)*/

	sort.Stable(sort.Reverse(sct))
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Sort After")
		spew.Dump(sct)
		fmt.Println(debugTraceEnd)*/
}

func (sc *society) calcAdjustmentFactor() {
	var scoreAccu float64
	count := len(sc.Members)
	for rank := range sc.Members {
		scoreAccu = calcScore(sc.Equality, float64(count), float64(rank))
	}
	sc.RreproduceArg.AdjustmentFactor = (1 / scoreAccu) / (float64(count) / float64(sc.Genesisarg.InitialPopulation))
}

type societySorterByFitness struct {
	toSort *society
}

func (sc *societySorterByFitness) Len() int {
	return len(sc.toSort.Members)
}

func (sc *societySorterByFitness) Swap(i, j int) {
	t := sc.toSort.Members[i]
	sc.toSort.Members[i] = sc.toSort.Members[j]
	sc.toSort.Members[j] = t
}

func (sc *societySorterByFitness) Less(i, j int) bool {
	return (*sc.toSort.Members[i]).getFitness() < (*sc.toSort.Members[j]).getFitness()
}

type societySorterGenid struct {
	toSort *society
}

func (sc *societySorterGenid) Len() int {
	return len(sc.toSort.Members)
}

func (sc *societySorterGenid) Swap(i, j int) {
	t := sc.toSort.Members[i]
	sc.toSort.Members[i] = sc.toSort.Members[j]
	sc.toSort.Members[j] = t
}

func (sc *societySorterGenid) Less(i, j int) bool {
	return (*sc.toSort.Members[i]).genid < (*sc.toSort.Members[j]).genid
}
