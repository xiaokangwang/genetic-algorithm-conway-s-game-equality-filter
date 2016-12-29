package main

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/davecgh/go-spew/spew"

	"golang.org/x/crypto/sha3"
)

type society struct {
	Members        [](*citizen)
	Achievement    int
	Generation     int
	Equality       float64
	Genesisarg     *genesisArg
	Fitnesscalcarg *fitnessCalcArg
	simarg         *SimGlobalArg
	mutationArg    *MutationArg
	reproduceArg   *ReproduceArg
}

type citizen struct {
	BelongTo *society
	Gene     *big.Int
	Fitness  int
	Parent   string
}

const fitnesso = "fitness %v %v\n"

func (ctz *citizen) getFitness() int {
	if ctz.Fitness == notCalc {
		ctz.Fitness = getFitness(ctz.Gene, ctz.BelongTo.Fitnesscalcarg)
	}
	fmt.Printf(fitnesso, ctz.getID(), ctz.Fitness)
	return ctz.Fitness
}

const tr = "trace/%v-%v"

func (ctz *citizen) getTrace() int {
	filename := fmt.Sprintf(tr, ctz.BelongTo.Generation, ctz.getID())
	ctz.Fitness = getFitnessTrace(ctz.Gene, ctz.BelongTo.Fitnesscalcarg, filename)

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
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.bph", s.Fitnesscalcarg.bph)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.bpw", s.Fitnesscalcarg.bpw)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.height", s.Fitnesscalcarg.height)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.width", s.Fitnesscalcarg.width)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.step", s.Fitnesscalcarg.step)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.offseth", s.Fitnesscalcarg.offseth)
	fmt.Printf("Metadata %v %v\n", "Fitnesscalcarg.offsetw", s.Fitnesscalcarg.offsetw)
	fmt.Printf("Metadata %v %v\n", "Genesisarg.initialPopulation", s.Genesisarg.initialPopulation)
	fmt.Printf("Metadata %v %v\n", "Genesisarg.density", s.Genesisarg.density)
	fmt.Printf("Metadata %v %v\n", "mutationArg.major.filp", s.mutationArg.major.filp)
	fmt.Printf("Metadata %v %v\n", "mutationArg.major.nega", s.mutationArg.major.nega)
	fmt.Printf("Metadata %v %v\n", "mutationArg.major.posi", s.mutationArg.major.posi)
	fmt.Printf("Metadata %v %v\n", "mutationArg.minor.filp", s.mutationArg.minor.filp)
	fmt.Printf("Metadata %v %v\n", "mutationArg.minor.nega", s.mutationArg.minor.nega)
	fmt.Printf("Metadata %v %v\n", "mutationArg.minor.posi", s.mutationArg.minor.posi)
	fmt.Printf("Metadata %v %v\n", "mutationArg.majorRate", s.mutationArg.majorRate)
	fmt.Printf("Metadata %v %v\n", "simarg.seed", s.simarg.seed)
	fmt.Printf("Metadata %v %v\n", "Equality", s.Equality)
	fmt.Printf("Metadata %v %v\n", "Ver", 1)
}

func (sc *society) sortCitizen() {
	sct := &societySorterByFitness{toSort: sc}

	fmt.Println(debugTraceStart)
	fmt.Println("Sort Before")
	spew.Dump(sct)
	fmt.Println(debugTraceEnd)

	sort.Sort(sort.Reverse(sct))

	fmt.Println(debugTraceStart)
	fmt.Println("Sort After")
	spew.Dump(sct)
	fmt.Println(debugTraceEnd)
}

func (sc *society) calcAdjustmentFactor() {
	var scoreAccu float64
	count := len(sc.Members)
	for rank := range sc.Members {
		scoreAccu = calcScore(sc.Equality, float64(count), float64(rank))
	}
	sc.reproduceArg.adjustmentFactor = (1 / scoreAccu) / (float64(count) / float64(sc.Genesisarg.initialPopulation))
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
