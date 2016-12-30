package main

import "fmt"

func main_disabled() {
	s := &society{}
	s.Fitnesscalcarg = &fitnessCalcArg{}
	s.Fitnesscalcarg.Bph = 64
	s.Fitnesscalcarg.Bpw = 64
	s.Fitnesscalcarg.Height = 128
	s.Fitnesscalcarg.Width = 128
	s.Fitnesscalcarg.Step = 100
	s.Fitnesscalcarg.Offseth = (s.Fitnesscalcarg.Height - s.Fitnesscalcarg.Bph) / 2
	s.Fitnesscalcarg.Offsetw = s.Fitnesscalcarg.Offseth
	s.Genesisarg = &genesisArg{}
	//s.Genesisarg.calcArg = *s.Fitnesscalcarg
	s.Genesisarg.InitialPopulation = 30
	s.Genesisarg.Density = 0.4
	s.MmutationArg = &MutationArg{}
	s.MmutationArg.Mmajor = MajorMutationArg{}
	s.MmutationArg.Mmajor.Filp = 0.05
	s.MmutationArg.Mmajor.Nega = 0.05
	s.MmutationArg.Mmajor.Posi = 0.05
	s.MmutationArg.Mminor = MiniorMutationArg{}
	s.MmutationArg.Mminor.Filp = 0.001
	s.MmutationArg.Mminor.Nega = 0.0005
	s.MmutationArg.Mminor.Posi = 0.0005
	s.MmutationArg.MmajorRate = 0.01
	s.Ssimarg = &SimGlobalArg{}
	s.Ssimarg.Seed = "FirstSeed2"
	s.Equality = 1.5
	s.RreproduceArg = &ReproduceArg{}
	s.RreproduceArg.AdjustmentFactor = 1
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
