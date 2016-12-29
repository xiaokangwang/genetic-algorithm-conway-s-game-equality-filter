package main

import "fmt"

func tryDensity(density float64) {
	fmt.Printf("subtest start\n")
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
	s.Genesisarg.density = density
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
	s.simarg.seed = "DensityPlot"
	s.Equality = 1.5
	s.reproduceArg = &ReproduceArg{}
	s.reproduceArg.adjustmentFactor = 1
	s.outputmeta()
	fmt.Printf("Genesis\n")
	s.Genesis(s.Genesisarg)
	fmt.Printf("Round\n")
	for i := 0; i < 1; i++ {
		s.legacy()
		s.outputMember()
		fmt.Printf("gen %v %v %v\n", s.Generation, s.Achievement, len(s.Members))
		s.generation()
	}
	fmt.Printf("subtest end\n")
}

func tryStep(steps int) {
	fmt.Printf("subtest start\n")
	s := &society{}
	s.Fitnesscalcarg = &fitnessCalcArg{}
	s.Fitnesscalcarg.bph = 64
	s.Fitnesscalcarg.bpw = 64
	s.Fitnesscalcarg.height = 128
	s.Fitnesscalcarg.width = 128
	s.Fitnesscalcarg.step = steps
	s.Fitnesscalcarg.offseth = (s.Fitnesscalcarg.height - s.Fitnesscalcarg.bph) / 2
	s.Fitnesscalcarg.offsetw = s.Fitnesscalcarg.offseth
	s.Genesisarg = &genesisArg{}
	s.Genesisarg.calcArg = *s.Fitnesscalcarg
	s.Genesisarg.initialPopulation = 30
	s.Genesisarg.density = 0.35
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
	s.simarg.seed = "StepPlot"
	s.Equality = 1.5
	s.reproduceArg = &ReproduceArg{}
	s.reproduceArg.adjustmentFactor = 1
	s.outputmeta()
	fmt.Printf("Genesis\n")
	s.Genesis(s.Genesisarg)
	fmt.Printf("Round\n")
	for i := 0; i < 1; i++ {
		s.legacy()
		s.outputMember()
		fmt.Printf("gen %v %v %v\n", s.Generation, s.Achievement, len(s.Members))
		s.generation()
	}
	fmt.Printf("subtest end\n")
}

func main() {
	for i := 1; i < 201; i = i + 20 {
		tryStep(i)
	}
}
