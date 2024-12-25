# Gata-Machine
Inspired by [ThePrimeagen/kata-machine](https://github.com/ThePrimeagen/kata-machine)

## Description
A script that generates algorithms for you to solve using Golang everyday.

## Installation
Clone the repository:
```bash
git clone https://github.com/VinewZ/gata-machine.git
```

## Usage
```bash
go run . generate
```
To generate a new batch of algorithms following the structure:
```
-_today
-day1
-day2
```
You should solve the algorithm in the _today directory.
The day[n] directy is maintained for reference.

## Testing
```bash
go test -v ./internal/__tests__/AlgoYouWantToTest_test.go
```
To test the algorithm you have solved.

or

```bash
go test -v ./internal/__tests__/
```
To test all the algorithms.

## Supported Algorithms
- WIP

## To-Do
- [ ]  Add tests for ALL algorithms
- [ ]  Add data structures
~maybe~
- [ ]  Create custom test runner
