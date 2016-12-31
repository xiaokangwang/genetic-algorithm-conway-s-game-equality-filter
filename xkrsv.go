package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
)

func Restore(seedname string, maxgen int) {
	s := &society{}
	s.Restore(seedname)
	s.outputmeta()
	if s.Ssimarg.Seed != seedname {
		panic("Seed not match")
	}
	if s.Generation == 0 {
		s.Genesis(s.Genesisarg)
		s.legacy()
		s.outputMember()
		s.CheckPoint()
	}
	var stop int32
	go func(intr *int32) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		_ = <-c
		atomic.AddInt32(intr, 1)
		_ = <-c
		atomic.AddInt32(intr, 1)
		_ = <-c
		os.Exit(-1)
	}(&stop)
	for s.Generation <= maxgen {
		s.generation()
		s.legacy()
		s.outputMember()
		fmt.Printf("gen %v %v %v\n", s.Generation, s.Achievement, len(s.Members))
		s.CheckPoint()
		sts := atomic.LoadInt32(&stop)
		if sts != 0 {
			os.Exit(0)
		}
	}
}

func main() {
	arg2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	Restore(os.Args[1], arg2)
}
