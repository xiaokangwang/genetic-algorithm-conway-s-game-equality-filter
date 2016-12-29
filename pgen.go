package main

import (
	"encoding/binary"
	"math/rand"

	"golang.org/x/crypto/sha3"
)

type PGSource struct {
	state sha3.ShakeHash
}

func (pg *PGSource) Int63() int64 {
	var ret int64
	binary.Read(pg.state, binary.BigEndian, &ret)
	ret = ret >> 1
	return ret
}

func (pg *PGSource) Seed(_ int64) {
}

func getPGSource(desc string) *PGSource {
	s := sha3.NewShake256()
	s.Write([]byte(desc))
	return &PGSource{state: s}
}

func getPGSourceRand(desc string) *rand.Rand {
	ps := getPGSource(desc)
	return rand.New(ps)
}
