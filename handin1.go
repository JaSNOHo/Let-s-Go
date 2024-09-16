/*
- The program ensures that the forks(represented by channels) are always
available in a non-blocking way.
	- The channel has a buffersize of 1, meaning it can only hold one message:
	true or waiting for a false.
- The goroutines constantly listens and handles the availability of the fork
asynchronously and randomised.
This combination ensures that philosophers can eventually get forks, eat, and
release the forks in a deadlock-free manner.
Meaning its non-blocking and prevents vircular waiting conditions which are the
main causes of deadlock.
*/
package main

func main() {
	table := make([]func(), 10)
	table[0] = hollenborg
	table[1] = fork
	table[2] = laksoe
	table[3] = pick
	table[4] = sandberg
	table[5] = prongs
	table[6] = descartes
	table[7] = trifork
	table[8] = niet
	table[9] = spork

	go hollenborg()
	go laksoe()
	go sandberg()
	go descartes()
	go niet()
	go fork()
	go pick()
	go prongs()
	go trifork()
	go pick()
	go spork()
}

func hollenborg() { //[0]
	var mealAmt int
	//request forks
	//print for every phil if they're eating or thinking
	//^ if !notEating : thinking
	if !table[9].isHeld && !table[1].isHeld {
		LOCK
		forks[9].isHeld = true
		forks[1].isHeld = true
		print("Hollenborg is eating")
		mealAmt++
		forks[9].isHeld = false
		forks[1].isHeld = false
		UNLOCK
	} else {
		print("Hollenborg is thinking")
	}
}

func laksoe() { //[2]
	//
}

func sandberg() { //[4]

}

func descartes() { //[6]

}

func niet() { //nietzsche [8]

}

func fork() { //[1]
	var isHeld bool

}

func prongs() { //[3]
	var isHeld bool
}

func trifork() { //[5]
	var isHeld bool
}

func pick() { //[7]
	var isHeld bool
}

func spork() { //[9]
	var isHeld bool = false
}

//arrays should be SEPERATE, or else fork and philos might change the array at
//the same time and fuck it UP

//use lock and unlock? lock forks when in use, unlock when not in use
//forks and philos need to communicate with channels
