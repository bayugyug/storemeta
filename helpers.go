package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"strconv"
)

func memDmp() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	m, _ := strconv.Atoi(strconv.FormatUint(mem.Alloc, 10))
	log.Println("Mem ", fmt.Sprintf("%.04f", float64(m)/(1024*1024)), " MB")
}

func statsDmp() {
	//sort
	keys := []string{}
	stats := pStats.getStatsList()
	for k := range stats {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// Loop over sorted key-value pairs.
	for i := range keys {
		kk := keys[i]
		vv := stats[kk]
		log.Println("Stats", kk, " -> ", vv)
	}
	log.Println("Elapsed:", pStats.elapsed())
}
