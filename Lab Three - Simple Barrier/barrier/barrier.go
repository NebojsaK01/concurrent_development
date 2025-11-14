//Barrier.go Template Code
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
// Created on 30/9/2024
// Modified by: Nebojsa Kukic
// Student Number: C00283550
// Issues:
// The barrier is not implemented! --solved--
// --------------------------------------------
// Description: The barrier has been implemented using mutexes and semaphores
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var barrierMutex sync.Mutex
var barrierCounter int
var barrierSem *semaphore.Weighted

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, totalRoutines int) /* bool not used*/ {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A

	barrierMutex.Lock()
	barrierCounter++
	if barrierCounter == totalRoutines {
		// Last goroutine to arrive - release all waiting goroutines
		barrierSem.Release(int64(totalRoutines))
	}
	barrierMutex.Unlock()

	// Wait at the barrier

	fmt.Println("PartB", goNum)
	wg.Done()
	// return true -> not used
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	barrierCounter = 0
	barrierSem = semaphore.NewWeighted(int64(totalRoutines))

	err := barrierSem.Acquire(ctx, int64(totalRoutines))
	if err != nil {
		return
	}

	for i := 0; i < totalRoutines; i++ { // Fixed: range over integer
		go doStuff(i, &wg, totalRoutines)
	}

	wg.Wait() //wait for everyone to finish before exiting

	fmt.Println("\nAll goroutines finished.")
}
