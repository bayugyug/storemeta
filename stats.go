package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

//StatsHelper - are holder of summary
type StatsHelper struct {
	locker sync.RWMutex
	stats  map[string]int
	t1     time.Time
	t2     time.Time
}

//StatsHelperNew - return new object
func StatsHelperNew() (n *StatsHelper) {
	//init
	return &StatsHelper{stats: make(map[string]int), t1: time.Now()}
}

func (s *StatsHelper) getStatsList() map[string]int {
	s.locker.Lock()
	defer s.locker.Unlock()
	if s.stats != nil {
		return s.stats
	}
	return nil
}

func (s *StatsHelper) setStats(prefix string) {
	if prefix == "" {
		return
	}
	s.locker.Lock()
	s.stats[prefix]++
	s.locker.Unlock()
}

func (s *StatsHelper) incStats(prefix string, more int) {
	if prefix == "" {
		return
	}
	s.locker.Lock()
	s.stats[prefix] += more
	s.locker.Unlock()
}

func (s *StatsHelper) getStats(prefix string) int {
	if prefix == "" {
		return 0
	}
	s.locker.Lock()
	m := s.stats[prefix]
	s.locker.Unlock()
	return m
}

func (s *StatsHelper) elapsedStart() {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.t1 = time.Now()
}

func (s *StatsHelper) elapsedDone() {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.t2 = time.Now()
}

func (s *StatsHelper) elapsed() string {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.t2 = time.Now()
	ts := s.t2.Sub(s.t1)
	return strings.TrimSpace(fmt.Sprintf("%02d:%02d:%02d", int(math.Mod(ts.Hours(), 12)), int(math.Mod(ts.Minutes(), 60)), int(math.Mod(ts.Seconds(), 60))))
}
