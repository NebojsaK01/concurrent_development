//Barrier.go Template Code - fixed
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

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Nebojsa Kukic
// Student Number: C00283550
// Description: The reusable barrier was implemented along with changing the mutex to an atomic var - int32.
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable --solved--
//2. Make it a reusable barrier --solved--
//-----

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Reusable barrier using atomic variables and unbuffered channel
func doStuff(goNum int, arrived *int32, max int, wg *sync.WaitGroup, theChan chan bool) /* bool -> not needed*/ {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// We wait here until everyone has completed part A
	// Use atomic operations instead of mutex
	newArrived := atomic.AddInt32(arrived, 1)

	if newArrived == int32(max) {
		// Last to arrive - signal others to go
		// Drain any remaining signals from previous use
		select {
		case <-theChan:
		default:
		}
		// Signal all waiting goroutines
		for i := 0; i < max-1; i++ {
			theChan <- true
		}
		// Reset for next use
		atomic.StoreInt32(arrived, 0)
	} else {
		// Not all here yet - wait until signaled
		<-theChan
	}

	fmt.Println("PartB", goNum)
	wg.Done()
	// return true -> not needed
}

func main() {
	totalRoutines := 10
	var arrived int32 // replace to int32, int64 is double the memory..
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	// Use unbuffered channel for signaling
	theChan := make(chan bool)

	// Create the goroutines
	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &arrived, totalRoutines, &wg, theChan)
	}

	wg.Wait() // Wait for everyone to finish before exiting

	// Showing the  reusability -> the barrier can be used again
	fmt.Println("\n-- Reusing Barrier --\n")

	var wg2 sync.WaitGroup
	wg2.Add(totalRoutines)
	atomic.StoreInt32(&arrived, 0) // Reset for reuse

	for i := 0; i < totalRoutines; i++ {
		go doStuff(i+totalRoutines, &arrived, totalRoutines, &wg2, theChan)
	}

	wg2.Wait()
	fmt.Println("\nAll goroutines finished.")
}
