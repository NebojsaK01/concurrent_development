# Documentation: Barrier Synchronization in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe


## Overview
This program implements a simple barrier in Go.  
A barrier makes all goroutines wait until everyone reaches the same point before continuing.


## How It Works
 - A mutex protects a counter that tracks how many goroutines have arrived.
 - When the last goroutine arrives, it releases all others using a weighted semaphore.
 - All goroutines then continue past the barrier and finish.

## Result
All "Part A" messages print first.  
Only after all goroutines reach the barrier do the "Part B" messages print.

This confirms the barrier works correctly.


