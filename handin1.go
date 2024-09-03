// must not lead to deadlock
// phil-forks are arranged around a round table alternately
// phils need 2 forks to eat, max 2 phils can eat at the same time
package main

func main() {
	//EACH fork must have a go routine
	//EACH philo must have a go routine
}

func hollenborg() {
	//request forks

	//print for every phil if they're eating or thinking
	//^ if !notEating : thinking
}

func laksoe() {

}

func sandberg() {

}

func decarte() {

}

func niet() { //nietzsche

}

func fork() {
	//fork-array? 5 spots, each spot is a corresponding table-placement?
	//or 10 spots, every second empty, used by a parallel philos array
}

func prongs() {

}

func trifork() {

}

func pick() {

}

func spork() {

}

//arrays should be SEPERATE, or else fork and philos might change the array at
//the same time and fuck it UP

//use lock and unlock? lock forks when in use, unlock when not in use
//forks and philos need to communicate with channels
