# Barrier Synchronization in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe

GitHub Repo: https://github.com/NebojsaK01/concurrent_development

The code implements a barrier using mutexes and semaphores
barrier ensures all goroutines all reach the same point before continuing.

### How to Run the Code:

#### run directly:
 - go run barrier.go

#### Build and Run:
 - go build -o barrier
 - ./barrier

#### Requirements:

- running on Go 1.25.3
-standard library packages:
  - "context"
  -  "fmt"
  -  "sync"
  -  "time"

- semaphore package: golang.org/x/sync/semaphore
