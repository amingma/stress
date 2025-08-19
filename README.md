# Stress Tester

A CLI tool built with Go and Cobra that helps competitive programmers validate their solutions by stress testing them against a known correct implementation.

## Description

This tool automates the process of stress testing competitive programming solutions Instead of manually creating test cases, the tool generates random inputs, runs both your solution and a reference "model" solution, and compares the outputs to identify discrepancies.

### Key Features

- **Automated Test Case Generation**: Generates random test cases based on your problem constraints
- **Comparative Testing**: Runs your solution against a trusted model solution
- **Detailed Reporting**: Provides comprehensive feedback on test results and performance
- **Flexible Configuration**: Customizable test parameters and input generation

### How It Works

1. **Setup**: You provide your solution and a model (correct) solution
2. **Generation**: The tool generates random test cases according to specified constraints
3. **Execution**: Both solutions are run on each generated test case
4. **Comparison**: Outputs are compared to detect differences
5. **Reporting**: Any discrepancies are reported with the failing test case

This approach is invaluable for finding corner cases that might not be covered in sample test cases, helping ensure your solution is robust before submission.

## Installation

### Prerequisites

- Go 1.24 or higher
- Compatible C++ compiler (only C++ solutions support as of 8/19)

### Install from Source

```bash
git clone https://github.com/amingma/stress
cd stress
go build 
go install
```

### Install with Go

```bash
go install github.com/amingma/stress
```

## Usage

### Basic Usage

```bash
stress test ./my_solution.cpp ./model_solution.cpp --cases 100 --type array
```

### Command Structure

```bash
stress test [path to your solution] [path to model solution] [flags]
```

### Available Commands

#### `test` - Execute Stress Test

Runs the stress testing process with your solution against the model solution.

**Optional Flags:**
- `--cases, -n`: Number of test cases to generate (default: 100)
- `--type, -t`: Type of problem (array, graph, tree)
- `--multi-query, -q`: Multiple test cases per test case (1 if yes, 0 if no)



