// must not lead to deadlock
// phil-forks are arranged around a round table alternately
// phils need 2 forks to eat, max 2 phils can eat at the same time
package main

func main() {
	go fork()
	//EACH fork must have a go routine
	go philo()
	//EACH philo must have a go routine
}

//do we need a func for each fork and each philo?

func fork() {
	//fork-array? 5 spots, each spot is a corresponding table-placement?
	//or 10 spots, every second empty, used by a parallel philos array
}

//arrays should be SEPERATE, or else fork and philos might change the array at
//the same time and fuck it UP
func philos() {
	//request forks

	//print for every phil if they're eating or thinking
	//^ if !notEating : thinking
}

//use lock and unlock? lock forks when in use, unlock when not in use
//forks and philos need to communicate with channels
