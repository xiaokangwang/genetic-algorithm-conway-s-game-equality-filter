package main

import "math/big"

type fitnessCalcArg struct {
	step                                      int
	width, height, offsetw, offseth, bpw, bph int
}

func getFitness(subject *big.Int, arg *fitnessCalcArg) int {
	field := loadFromBitwiseRepresention(arg.width, arg.height, arg.offsetw, arg.offseth, arg.bpw, arg.bph, subject)
	for cs := 0; cs < arg.step; cs++ {
		field.nextRound()
	}
	field.Finialize()
	return field.CountLife()
}

func getFitnessTrace(subject *big.Int, arg *fitnessCalcArg, trace string) int {
	field := loadFromBitwiseRepresention(arg.width, arg.height, arg.offsetw, arg.offseth, arg.bpw, arg.bph, subject)
	field.setTrace(trace)
	for cs := 0; cs < arg.step; cs++ {
		field.nextRound()
	}
	field.Finialize()
	return field.CountLife()
}
