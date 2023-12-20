# Go Grep Implementations

## Overview
This repository contains a lightweight demonstration of Grep in Go, showcasing five different implementations: `minimal`, `bad`, `ok`, `good`, and `great`. Each implementation varies in complexity and performance, providing a unique perspective on how the Grep functionality can be achieved in Go.

## Implementations
- **Minimal:** A very basic implementation with limited functionality.
- **Bad:** An example of how not to implement Grep, useful for learning purposes.
- **Ok:** A decent implementation with some improvements over the basic version.
- **Good:** A well-rounded implementation with better readability.
- **Great:** The most advanced implementation, using concurrency.

## Usage
To run any of these implementations, use the Go command-line tool with appropriate flags. The available options for each implementation can be viewed using the flag library.

### Example
An example of running the program with the `great` implementation is shown below:

```bash
go run grep.go -n -i -l -g great "SON OF PELEUS" texts/frankenstein.txt texts/iliad.txt
