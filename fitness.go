package main

import "math/big"

type fitnessCalcArg struct {
	Step                                      int
	Width, Height, Offsetw, Offseth, Bpw, Bph int
}

func getFitness(subject *big.Int, arg *fitnessCalcArg) int {
	field := loadFromBitwiseRepresention(arg.Width, arg.Height, arg.Offsetw, arg.Offseth, arg.Bpw, arg.Bph, subject)
	for cs := 0; cs < arg.Step; cs++ {
		field.nextRound()
	}
	field.Finialize()
	return field.CountLife()
}

func getFitnessTrace(subject *big.Int, arg *fitnessCalcArg, trace string) int {
	field := loadFromBitwiseRepresention(arg.Width, arg.Height, arg.Offsetw, arg.Offseth, arg.Bpw, arg.Bph, subject)
	field.setTrace(trace)
	for cs := 0; cs < arg.Step; cs++ {
		field.nextRound()
	}
	field.Finialize()
	return field.CountLife()
}
