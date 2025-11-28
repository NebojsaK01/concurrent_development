# Dining Philosophers in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe

GitHub Repo: https://github.com/NebojsaK01/concurrent_development

## Overview

This lab simulates the Dining Philosophers problem in Go.
Philosophers alternate between thinking and eating, and a host semaphore 
prevents deadlock by allowing only 4 of 5 philosophers to pick up forks simultaneously.

## How to Run the Code:

#### Run directly:

go run dinPhil(1).go

#### Build and Run:

go build -o dinPhil(1)
./dinPhil(1)

#### Requirements:

-Running on Go 1.25.3
-Standard library packages:
  - fmt
  - sync
  - time
  - math/rand

