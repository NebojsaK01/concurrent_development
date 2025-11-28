# Documentation: Reusable Barrier in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe


## Overview

This program implements a reusable barrier in Go.
All goroutines wait at the barrier until everyone reaches the same point before continuing.
The barrier can be used multiple times.

## How It Works

- An atomic counter (int32) tracks how many goroutines have arrived.
- An unbuffered channel signals goroutines to continue once the last goroutine arrives.
- The last goroutine releases all others and resets the counter for reuse.
- All goroutines then continue past the barrier and finish.

## Result
All "Part A" messages print first.

Only after all goroutines reach the barrier do the "Part B" messages print.
The program demonstrates reusability by running a second round of goroutines.


