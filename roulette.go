package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	capacityMin = 1
	capacityMax = 7 // Min-Max
)

// Candidate is an int alias for roulette
type Candidate int

// randInt returns an int between min+(rand(max-min))
func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generator generates candidates while signalled to do so
func Generator(pending chan<- Candidate, ready <-chan bool) {
	for {
		<-ready // blocking read
		nc := Candidate(randInt(capacityMin, capacityMax))
		log.Println("new candidate:", int(nc))
		pending <- nc
	}
}

// Lucky determines luckiness of the candidate with equality
func (c Candidate) Lucky() bool {
	return c == 6
}

func main() {
	// seed rand
	rand.Seed(time.Now().UnixNano())

	candidates := make(chan Candidate, 1)
	ready := make(chan bool, 1)

	// Prime the generator
	ready <- true

	// Generator on the candidates should block.. until,..
	go Generator(candidates, ready)

	for {
		// Feeling lucky, punk?
		if c := <-candidates; c.Lucky() {
			log.Println("lucky candidate:", c)
			break
		}
		// Oh well, reload
		ready <- true
	}
}
