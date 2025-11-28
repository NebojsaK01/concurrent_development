# Documentation: Dining Philosophers in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe

## Overview

This program simulates the Dining Philosophers problem in Go.
Philosophers alternate between thinking and eating.
A host semaphore ensures only 4 out of 5 philosophers can pick up forks at once, preventing deadlock.

## How It Works

- Each philosopher is a goroutine that cycles through thinking and eating.
- Forks are represented as buffered channels; picking up a fork means sending into the channel.
- A host semaphore limits how many philosophers can attempt to pick up forks simultaneously.
- Philosophers request permission from the host, pick up forks, eat, then put down forks and release the host slot.

## Result

- The host ensures no deadlock occurs (circular wait deadlock).
- All philosophers complete their eating cycles safely.
- This confirms correct synchronization and demonstrates deadlock prevention.



