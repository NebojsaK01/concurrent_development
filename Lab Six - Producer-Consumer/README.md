# Producer Consumer in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe

GitHub Repo: https://github.com/NebojsaK01/concurrent_development

## Overview

This lab implements the Producer-Consumer problem in Go using a thread-safe buffer with limited capacity.
The solution uses mutexes and condition variables to synchronize access between producers and consumers,
preventing race conditions and ensuring efficient resource utilization.

## How to Run the Code:

#### Run directly:

go run producerConsumer.go

#### Build and Run:

go build -o producerConsumer
./producerConsumer

#### Requirements:

- Running on Go 1.25.3
- Standard library packages:
    - fmt
    - sync
    - time