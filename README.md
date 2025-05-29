# Hello Bazel 👋

A collection of simple Bazel examples demonstrating how to build projects in different programming languages.

## 📦 Contents

1. **Go Example** (`//hello-go:hello-go`)
   - Basic Go program
   - Shows Go integration with Bazel

3. **Python Example** (`//hello-py:hello`)
   - Python binary and library
   - Demonstrates PyBinary rule

## 🚀 Getting Started

### Prerequisites
- [Bazel](https://bazel.build/) installed (version 6.x or later recommended)
- For specific examples:
  - Go toolchain (for Go examples)
  - Python (for Python examples)

### Building Examples

To build any target, run:

```bash
bazel build //hello-py:hello

bazel build //hello-go:hello-go
```

To run a target directly:

```bash
bazel run //hello-go:gazelle
bazel run //hello-go:hello-go

bazel run //hello-py:hello
```
