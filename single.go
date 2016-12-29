package main

import "fmt"

func main_disabled() {
	s := &society{}
	s.Fitnesscalcarg = &fitnessCalcArg{}
	s.Fitnesscalcarg.bph = 64
	s.Fitnesscalcarg.bpw = 64
	s.Fitnesscalcarg.height = 128
	s.Fitnesscalcarg.width = 128
	s.Fitnesscalcarg.step = 100
	s.Fitnesscalcarg.offseth = (s.Fitnesscalcarg.height - s.Fitnesscalcarg.bph) / 2
	s.Fitnesscalcarg.offsetw = s.Fitnesscalcarg.offseth
	s.Genesisarg = &genesisArg{}
	s.Genesisarg.calcArg = *s.Fitnesscalcarg
	s.Genesisarg.initialPopulation = 30
	s.Genesisarg.density = 0.4
	s.mutationArg = &MutationArg{}
	s.mutationArg.major = majorMutationArg{}
	s.mutationArg.major.filp = 0.05
	s.mutationArg.major.nega = 0.05
	s.mutationArg.major.posi = 0.05
	s.mutationArg.minor = miniorMutationArg{}
	s.mutationArg.minor.filp = 0.001
	s.mutationArg.minor.nega = 0.0005
	s.mutationArg.minor.posi = 0.0005
	s.mutationArg.majorRate = 0.01
	s.simarg = &SimGlobalArg{}
	s.simarg.seed = "FirstSeed2"
	s.Equality = 1.5
	s.reproduceArg = &ReproduceArg{}
	s.reproduceArg.adjustmentFactor = 1
	s.outputmeta()
	fmt.Printf("Genesis\n")
	s.Genesis(s.Genesisarg)
	fmt.Printf("Round\n")
	for i := 0; i < 1000; i++ {
		s.legacy()
		s.outputMember()
		fmt.Printf("gen %v %v %v\n", s.Generation, s.Achievement, len(s.Members))
		s.generation()
	}
}
