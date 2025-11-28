# Reusable Barrier in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe

GitHub Repo: https://github.com/NebojsaK01/concurrent_development

This lab implements a reusable barrier using atomic variables and an unbuffered channel.  
The barrier ensures that all goroutines reach a synchronization point before any continue, 
and can be reused multiple times.

### How to Run the Code:

#### run directly:
 - go run barrier.go

#### Build and Run:
 - go build -o barrier
 - ./barrier

#### Requirements:

- running on Go 1.25.3
- standard library packages:
  - fmt
  - sync
  - sync/atomic
  - time

