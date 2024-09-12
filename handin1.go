// must not lead to deadlock
// phil-forks are arranged around a round table alternately
// phils need 2 forks to eat, max 2 phils can eat at the same time

package main

import (
	"fmt"
	"math/rand"
	"time"
)
//Fork struct represents the forks, communicating if its available by a channel
type Fork struct {
	id int
	ch chan bool
}
//Philosher struct represents the philosopher, keeps tracks of the forks
type Philosopher struct {
	id int
	leftFork *Fork
	rightFork *Fork
	mealsAmt int
}
//make philos w the feilds from Philosopher?
func hollenborg(left, right *Fork, done chan bool) {
	philosopher := Philosopher{id: 1, leftFork: left, rightFork: right}
	philosopher.dine(done)
}

func laksoe(left, right *Fork, done chan bool) {
	philosopher := Philosopher{id: 2, leftFork: left, rightFork: right}
	philosopher.dine(done)
}

func sandberg(left, right *Fork, done chan bool) {
	philosopher := Philosopher{id: 3, leftFork: left, rightFork: right}
	philosopher.dine(done)
}

func descartes(left, right *Fork, done chan bool) {
	philosopher := Philosopher{id: 4, leftFork: left, rightFork: right}
	philosopher.dine(done)
}

func niet(left, right *Fork, done chan bool) {
	philosopher := Philosopher{id: 5, leftFork: left, rightFork: right}
	philosopher.dine(done)
}

// Philosopher's dine function handles the philosopher's behavior.
func (p *Philosopher) dine(done chan bool) {
	for p.mealsAmt < 3 {
		// Philosopher is thinking
		fmt.Printf("Philosopher %d is thinking.\n", p.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Request the left fork
		<-p.leftFork.ch
		fmt.Printf("Philosopher %d picked up fork %d (left).\n", p.id, p.leftFork.id)

		// Request the right fork
		<-p.rightFork.ch
		fmt.Printf("Philosopher %d picked up fork %d (right).\n", p.id, p.rightFork.id)

		// Philosopher is eating
		fmt.Printf("Philosopher %d is eating.\n", p.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		p.mealsAmt++

		// Put down the right fork
		p.rightFork.ch <- true
		fmt.Printf("Philosopher %d put down fork %d (right).\n", p.id, p.rightFork.id)

		// Put down the left fork
		p.leftFork.ch <- true
		fmt.Printf("Philosopher %d put down fork %d (left).\n", p.id, p.leftFork.id)
	}
	done <- true
}

// Forks are goroutines that listen for release signals from philosophers.
func fork(f *Fork) {
	for {
		f.ch <- true // Signal that the fork is available
		<-f.ch       // Wait for a philosopher to take it
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create 5 forks and their channels
	forks := make([]*Fork, 5)
	for i := 0; i < 5; i++ {
		forks[i] = &Fork{id: i + 1, ch: make(chan bool, 1)}
		go fork(forks[i])
	}

	// Channel to indicate philosophers have finished eating
	done := make(chan bool)

	// Create and start the philosophers (mapping to your named functions)
	go hollenborg(forks[0], forks[1], done)
	go laksoe(forks[1], forks[2], done)
	go sandberg(forks[2], forks[3], done)
	go descartes(forks[3], forks[4], done)
	go niet(forks[4], forks[0], done)

	// Wait for all philosophers to finish
	for i := 0; i < 5; i++ {
		<-done
	}

	fmt.Println("All philosophers have eaten at least 3 times.")
}

