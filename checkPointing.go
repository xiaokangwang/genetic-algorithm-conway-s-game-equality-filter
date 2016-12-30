package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (sc *society) CheckPoint() {
	detailfoldername := fmt.Sprintf("checkpoint/%v/%v/", sc.Ssimarg.Seed, sc.Generation)
	_ := os.MkdirAll(detailfoldername, 0777)
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Checkpoint cp")
		fmt.Println(debugTraceEnd)
		if err != nil {
			fmt.Println("Checkpoint Failed! MkdirAll", err)
		}
	*/
	file, _ := os.Create(detailfoldername + "socity.json")
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Checkpoint cp")
		if err != nil {
			fmt.Println("Checkpoint Failed! FileOpen", err)
		}
		fmt.Println(debugTraceEnd)*/
	jsonw := json.NewEncoder(file)
	/*
		fmt.Println(debugTraceEnd)*/
	jsonw.SetIndent("", "    ")
	_ = jsonw.Encode(sc)
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Checkpoint cp")
		if err != nil {
			fmt.Println("Checkpoint Failed! Encode", err)
		}
		fmt.Println(debugTraceEnd)
	*/
	file.Close()
	currentSymname := fmt.Sprintf("checkpoint/%v/current", sc.Ssimarg.Seed)
	os.Remove(currentSymname)
	_ = os.Symlink("../../"+detailfoldername+"/socity.json", currentSymname)
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Checkpoint cp")
		if err != nil {
			fmt.Println("Checkpoint Failed! Symlink", err)
		}
		fmt.Println(debugTraceEnd)*/
}

func (sc *society) Restore(seedname string) {
	currentSymname := fmt.Sprintf("checkpoint/%v/current", seedname)
	file, _ := os.Open(currentSymname)
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Checkpoint rs")
		if err != nil {
			fmt.Println("Restore Failed! Open", err)
		}
		fmt.Println(debugTraceEnd)*/
	dec := json.NewDecoder(file)
	_ = dec.Decode(sc)
	/*
		fmt.Println(debugTraceStart)
		fmt.Println("Checkpoint rs")
		if err != nil {
			fmt.Println("Restore Failed! Decode", err)
		}
		fmt.Println(debugTraceEnd)*/
	file.Close()
	if sc.Members != nil {
		for _, mem := range sc.Members {
			mem.belongTo = sc
		}
	}
}
