# Car Builder Pattern

This repository implements the Creational Builder design pattern to construct complex `Car` objects step-by-step[cite: 2]. By separating the construction process from the object's representation, the same process is used to create different variations of an object[cite: 2].

## Features

* **Fluent API**: Builder methods return the builder instance to allow continuous method chaining[cite: 2].
* **Multiple Representations**: The identical construction steps produce two different representations—a structured `Car` object and a command-line string array[cite: 1].
* **Director Class**: Manages the construction process to easily reuse specific configurations, such as `MakeSportsCar` and `MakeCityCar`[cite: 2].
* **State Validation**: The final retrieval method validates the state and triggers an error on invalid or incomplete input[cite: 1].

## Quick Start

Execute the main file to see the Director build both the Object representation and the Command representation:

```bash
go run main.go
