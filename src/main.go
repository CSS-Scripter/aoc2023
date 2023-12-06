package main

import (
	"aoc2023/src/d01"
	"aoc2023/src/d02"
	"aoc2023/src/d03"
	"aoc2023/src/d04"
	"aoc2023/src/d05"
	"aoc2023/src/d06"
	"fmt"
	"sort"
	"time"

	"github.com/charmbracelet/log"
)

type Day struct {
	Day      int
	MainFunc func()
}

var days = []Day{
	{1, d01.Main},
	{2, d02.Main},
	{3, d03.Main},
	{4, d04.Main},
	{5, d05.Main},
	{6, d06.Main},
}

func main() {
	sort.Slice(days[:], func(i, j int) bool {
		return days[i].Day < days[j].Day
	})

	totalStart := time.Now()
	log.Default().SetLevel(log.DebugLevel)
	var start time.Time
	var elapsed time.Duration
	for _, day := range days {
		log.Default().Info(fmt.Sprintf("starting day %d", day.Day))
		start = time.Now()
		day.MainFunc()
		elapsed = time.Since(start)
		log.Default().Info(fmt.Sprintf("day %d calculated in %s\n", day.Day, elapsed))
	}
	totalDuration := time.Since(totalStart)
	log.Info(fmt.Sprintf("total duration: %s", totalDuration))
}
