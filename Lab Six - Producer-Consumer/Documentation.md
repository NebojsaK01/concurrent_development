# Documentation: Producer Consumer in Go

Student Name: Nebojsa Kukic

Student Number: C00283550

Lecturer Name: Joseph Kehoe

## Overview

This program simulates the Producer-Consumer problem in Go using a thread-safe buffer with limited capacity. 
Producers create items and add them to the buffer, while consumers remove items and process them. 
The solution ensures proper synchronization and prevents race conditions.

## How It Works

- The SafeBuffer uses a mutex to protect concurrent access to the buffer.
- Two condition variables (items and spaces) handle synchronization between producers and consumers.
- Producers wait when the buffer is full (spaces condition).
- Consumers wait when the buffer is empty (items condition).
- When a producer adds an item, it signals the items condition to wake up waiting consumers.
- When a consumer removes an item, it signals the spaces condition to wake up waiting producers.

## Result

- The buffer correctly limits capacity and prevents overflow.
- Producers and consumers synchronize properly without race conditions.
- All threads complete their work without deadlocks or starvation.
- This demonstrates effective use of mutexes and condition variables for thread synchronization.