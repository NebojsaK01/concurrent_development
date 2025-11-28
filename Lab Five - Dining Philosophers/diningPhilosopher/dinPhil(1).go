// dinPhil(1).go Template Code
// Author: Joseph Kehoe

//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on Created: 21/10/24
// Modified by: Nebojsa Kukic
// Student Number: C00283550
// Issues:
// All dining phils could pick a left fork at the same time: starve right forks.
// Therefore a circular wait deadlock occurs. - fixed.
// --------------------------------------------

// Description: Now it cannot deadlock as the host semaphore allows 4 out of 5 dining philosophers
// to pick up a fork at once.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// Added host semaphore to limit concurrent fork pickups
func getForks(index int, forks map[int]chan bool, host chan bool) {
	host <- true               // ask host for permission to pick up forks
	forks[index] <- true       // Pick up left fork
	forks[(index+1)%5] <- true // Pick up right fork
}

func putForks(index int, forks map[int]chan bool, host chan bool) {
	<-forks[index]       // Put down left fork
	<-forks[(index+1)%5] // put down right fork
	<-host               // Tell host we're done (free up a slot)
}

// Added host channel parameter
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, host chan bool) {
	defer wg.Done() // defer to ensure to always call wg.Done()

	for i := 0; i < 3; i++ { // Limited to 3 cycles to prevent infinite loop
		think(index)
		getForks(index, forks, host) // Pass host channel
		eat(index)
		putForks(index, forks, host) // Pass host channel
	}
	fmt.Printf("Phil %d finished eating\n", index)
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	// Added host semaphore - only 4 philosophers can eat at once
	host := make(chan bool, 4) // Prevents deadlock!

	forks := make(map[int]chan bool)
	for k := 0; k < philCount; k++ { // while to a for loop
		forks[k] = make(chan bool, 1) // Buffered channel = fork available
	}

	fmt.Println("Started the Dining Philosophers:")

	for N := 0; N < philCount; N++ { // while to a for loop
		go doPhilStuff(N, &wg, forks, host) // pass host to philosophers
	}

	wg.Wait() //wait here until everyone is done
	fmt.Println("\nAll philosophers have finished dining.")
} // main
