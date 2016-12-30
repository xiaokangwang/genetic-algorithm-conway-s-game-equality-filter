package main

import "fmt"

func tryDensity(density float64) {
	fmt.Printf("subtest start\n")
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
	s.Genesisarg.Density = density
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
	s.Ssimarg.Seed = "DensityPlot"
	s.Equality = 1.5
	s.RreproduceArg = &ReproduceArg{}
	s.RreproduceArg.AdjustmentFactor = 1
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
	s.Fitnesscalcarg.Bph = 64
	s.Fitnesscalcarg.Bpw = 64
	s.Fitnesscalcarg.Height = 128
	s.Fitnesscalcarg.Width = 128
	s.Fitnesscalcarg.Step = steps
	s.Fitnesscalcarg.Offseth = (s.Fitnesscalcarg.Height - s.Fitnesscalcarg.Bph) / 2
	s.Fitnesscalcarg.Offsetw = s.Fitnesscalcarg.Offseth
	s.Genesisarg = &genesisArg{}
	//s.Genesisarg.calcArg = *s.Fitnesscalcarg
	s.Genesisarg.InitialPopulation = 30
	s.Genesisarg.Density = 0.35
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
	s.Ssimarg.Seed = "StepPlot"
	s.Equality = 1.5
	s.RreproduceArg = &ReproduceArg{}
	s.RreproduceArg.AdjustmentFactor = 1
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

func tryEqu(tequality float64) {
	fmt.Printf("subtest start\n")
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
	s.Genesisarg.InitialPopulation = 256
	s.Genesisarg.Density = 0.35
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
	s.Ssimarg.Seed = "StepPlot"
	s.Equality = tequality
	s.RreproduceArg = &ReproduceArg{}
	s.RreproduceArg.AdjustmentFactor = 1
	s.outputmeta()
	fmt.Printf("Genesis\n")
	s.Genesis(s.Genesisarg)
	fmt.Printf("Round\n")
	for i := 0; i < 320; i++ {
		s.legacy()
		s.outputMember()
		fmt.Printf("gen %v %v %v\n", s.Generation, s.Achievement, len(s.Members))
		s.generation()
	}
	fmt.Printf("subtest end\n")
}

func tryInt(tequality float64) {
	fmt.Printf("subtest start\n")
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
	s.Genesisarg.Density = 0.35
	s.MmutationArg = &MutationArg{}
	s.MmutationArg.Mmajor = MajorMutationArg{}
	s.MmutationArg.Mmajor.Filp = 0.05
	s.MmutationArg.Mmajor.Nega = 0.05
	s.MmutationArg.Mmajor.Posi = 0.05
	s.MmutationArg.Mminor = MiniorMutationArg{}
	s.MmutationArg.Mminor.Filp = 0.0001
	s.MmutationArg.Mminor.Nega = 0.00005
	s.MmutationArg.Mminor.Posi = 0.00005
	s.MmutationArg.MmajorRate = 0.1
	s.Ssimarg = &SimGlobalArg{}
	s.Ssimarg.Seed = "IntTest"
	s.Equality = tequality
	s.RreproduceArg = &ReproduceArg{}
	s.RreproduceArg.AdjustmentFactor = 1
	s.outputmeta()
	fmt.Printf("Genesis\n")
	s.Genesis(s.Genesisarg)
	fmt.Printf("Round\n")
	for i := 0; i < 320; i++ {
		s.legacy()
		s.outputMember()
		fmt.Printf("gen %v %v %v\n", s.Generation, s.Achievement, len(s.Members))
		s.CheckPoint()
		s.generation()
	}
	fmt.Printf("subtest end\n")
}

func main_() {
	tryInt(1.5)
}
