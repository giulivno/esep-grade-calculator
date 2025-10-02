# ESEP Grade Calculator

This library is a simple utility to calculate final course grades based on weighted averages.

- **Assignments** – 50%  
- **Exams** – 35%  
- **Essays** – 15%

Runtime requirements:  
Golang 1.24+

---

## Setup
Clone your fork of the repository and change into the project directory:

git clone <your-fork-url>
cd grade-calculator/

---

## Run unit tests

go test -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt

---

## Project structure

go.mod # Go module definition
grade_calculator.go # Core library implementation
grade_calculator_test.go # Unit tests